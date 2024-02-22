package listener

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type snsSubscriptionMessage struct {
	Type              string                       `json:"Type"`
	MessageId         string                       `json:"MessageId"`
	TopicArn          string                       `json:"TopicArn"`
	Subject           string                       `json:"Subject"`
	Message           string                       `json:"Message"`
	Timestamp         string                       `json:"Timestamp"`
	SignatureVersion  string                       `json:"SignatureVersion"`
	Signature         string                       `json:"Signature"`
	SigningCertURL    string                       `json:"SigningCertURL"`
	UnsubscribeURL    string                       `json:"UnsubscribeURL"`
	MessageAttributes map[string]map[string]string `json:"MessageAttributes"`
}

type snsListener struct {
	lvl              zerolog.Level
	log              zerolog.Logger
	hdl              Handler
	protocol         string
	callbackListener net.Listener
	callbackUrl      string
	client           *sns.Client
}

// A listener the listens for SNS messages via public endpoint
// Not currently used cause I dont really like this approach.
// Maybe be fine to require this to poll messages from an SQS queue instead and register SQS as the reciever of SNS messages to avoid having to run another public endpoint.
func NewSnsListener(client *sns.Client, hdl Handler, opts ...ListenerOption) Listener {
	l := &snsListener{
		lvl:    zerolog.InfoLevel,
		log:    zerolog.Nop(),
		hdl:    hdl,
		client: client,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *snsListener) SetLogLevel(level zerolog.Level) {
	l.lvl = level
}

func (l *snsListener) SetLogger(logger zerolog.Logger) {
	l.log = logger
}

func (l *snsListener) Subscribe(ctx context.Context, topic string) error {
	go func() {
		http.Serve(l.callbackListener, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var msg snsSubscriptionMessage
			if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
				l.log.Error().Err(err).Msg("Failed to decode message")
				http.Error(w, "Failed to decode message", http.StatusBadRequest)
				return
			}

			lctx := logger.InitContextLogger(context.Background(), l.lvl)

			m, err := notification.DeserializeExecutionMessage([]byte(msg.Message))
			if err != nil {
				log.Error().Err(err).Msg("Failed to deserialize message")
				return
			}

			l.hdl.Handle(lctx, m)

			w.WriteHeader(http.StatusOK)
		}))

	}()

	go func() {
		<-ctx.Done()
		log.Debug().Msg("Shutting down server...")
	}()

	l.client.Subscribe(ctx, &sns.SubscribeInput{
		TopicArn: &topic,
		Protocol: aws.String(l.protocol),
		Endpoint: aws.String(l.callbackUrl),
	})

	return nil
}

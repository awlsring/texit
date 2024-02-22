package listener

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type mqttListener struct {
	lvl zerolog.Level
	log zerolog.Logger
	hdl Handler
	m   mqtt.Client
}

func (l *mqttListener) publishHandler(client mqtt.Client, msg mqtt.Message) {
	ctx := logger.InitContextLogger(context.Background(), l.lvl)
	log.Debug().Msg("Deserializing message")
	m, err := notification.DeserializeExecutionMessage(msg.Payload())
	if err != nil {
		log.Error().Err(err).Msg("Failed to deserialize message")
		return
	}
	log.Debug().Interface("message", m).Msg("Deserialized message")
	l.hdl.Handle(ctx, m)
}

func (l *mqttListener) connectHandler(mqtt.Client) {
	log.Info().Msg("Connected to broker")
}

func (l *mqttListener) connectDropHandler(mqtt.Client, error) {
	log.Error().Msg("Connection dropped")
}

func (l *mqttListener) Subscribe(ctx context.Context, t string) error {
	if token := l.m.Subscribe(t, 0, nil); token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("Failed to subscribe")
		return token.Error()
	}
	return nil
}

func (l *mqttListener) SetLogLevel(level zerolog.Level) {
	l.lvl = level
}

func (l *mqttListener) SetLogger(logger zerolog.Logger) {
	l.log = logger
}

func NewMqttListener(addr string, hdl Handler, opts ...ListenerOption) (Listener, error) {
	l := &mqttListener{
		lvl: zerolog.InfoLevel,
		log: log.Logger,
		hdl: hdl,
	}

	for _, o := range opts {
		o(l)
	}

	o := mqtt.NewClientOptions()
	o.AddBroker(addr)
	o.SetDefaultPublishHandler(l.publishHandler)
	o.SetOnConnectHandler(l.connectHandler)
	o.SetConnectionLostHandler(l.connectDropHandler)
	o.SetAutoReconnect(true)
	c := mqtt.NewClient(o)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	l.m = c

	return l, nil
}

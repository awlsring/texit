package sns_notification_gateway

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SnsGateway struct {
	client *sns.Client
	topic  string
}

func New(topic string, client *sns.Client) gateway.Notification {
	return &SnsGateway{
		topic:  topic,
		client: client,
	}
}

func (g *SnsGateway) Endpoint() string {
	return g.topic
}

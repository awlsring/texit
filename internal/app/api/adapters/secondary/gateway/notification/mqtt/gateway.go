package mqtt_notification_gateway

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttGateway struct {
	topic  string
	client mqtt.Client
}

func New(topic string, client mqtt.Client) gateway.Notification {
	return &MqttGateway{
		topic:  topic,
		client: client,
	}
}

func (g *MqttGateway) Endpoint() string {
	return g.topic
}

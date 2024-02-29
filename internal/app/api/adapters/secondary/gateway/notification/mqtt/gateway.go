package mqtt_notification_gateway

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttGateway struct {
	topic  string
	broker string
	client mqtt.Client
}

func New(broker, topic string, client mqtt.Client) gateway.Notification {
	return &MqttGateway{
		broker: broker,
		topic:  topic,
		client: client,
	}
}

func (g *MqttGateway) Type() notification.TopicType {
	return notification.TopicTypeMqtt
}

func (g *MqttGateway) Endpoint() notification.Endpoint {
	end := fmt.Sprintf("%s/%s", g.broker, g.topic)
	return notification.Endpoint(end)
}

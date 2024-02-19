package mqtt_notification_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
)

func (g *MqttGateway) SendMessage(ctx context.Context, message string) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Sending message to MQTT")

	t := g.client.Publish(g.topic, 0, false, message)
	if t.Wait() && t.Error() != nil {
		log.Error().Err(t.Error()).Msg("Failed to send message to MQTT")
		return t.Error()
	}

	log.Debug().Msg("Message sent")
	return nil
}

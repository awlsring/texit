package mqtt_notification_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
)

func (g *MqttGateway) SendMessage(ctx context.Context, message string) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Sending message to MQTT")

	t := g.client.Publish(g.topic, 0, false, message)
	go func() {
		<-t.Done()
		if err := t.Error(); err != nil {
			log.Err(err).Msg("Error sending message")
		} else {
			log.Debug().Msg("Message sent")
		}
	}()

	log.Debug().Msg("Message sent")
	return nil
}

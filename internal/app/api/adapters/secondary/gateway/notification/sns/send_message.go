package sns_notification_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func (g *SnsGateway) SendMessage(ctx context.Context, message string) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Sending message to SNS")

	_, err := g.client.Publish(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: &g.topic,
	})
	if err != nil {
		log.Err(err).Msg("Error sending message")
		return err
	}

	log.Debug().Msg("Message sent")
	return nil
}

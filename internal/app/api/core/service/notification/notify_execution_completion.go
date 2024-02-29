package notification

import (
	"context"
	"sync"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) NotifyExecutionCompletion(ctx context.Context, e workflow.ExecutionIdentifier, w workflow.WorkflowName, status workflow.Status, results workflow.ExecutionResult) error {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Notifying workflow end for %s", e.String())

	res, err := results.Serialize()
	if err != nil {
		log.Error().Err(err).Msg("Failed to serialize results")
		return err
	}

	msg := "Workflow ended with status " + status.String()
	message := notification.NewExecutionMessage(e, w, status, msg, res.String())

	log.Debug().Msg("Serializing message")
	smsg, err := message.Serialize()
	if err != nil {
		log.Error().Err(err).Msg("Failed to serialize message")
		return err
	}

	log.Debug().Msg("Sending notification to configured gateways")
	var wg sync.WaitGroup
	for _, gw := range s.gateways {
		wg.Add(1)
		go func(gw gateway.Notification) {
			defer wg.Done()
			err := gw.SendMessage(ctx, smsg)
			if err != nil {
				log.Warn().Err(err).Str("endpoint", gw.Endpoint().String()).Msg("Failed to send notification")
			}
		}(gw)
	}
	wg.Wait()

	log.Debug().Msg("Notification sent")
	return nil
}

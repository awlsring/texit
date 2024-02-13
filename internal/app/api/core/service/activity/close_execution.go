package activity

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) CloseExecution(ctx context.Context, exId workflow.ExecutionIdentifier, status workflow.Status, results workflow.ExecutionResult) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Closing execution")

	log.Debug().Msgf("Serializing output")
	res, err := results.Serialize()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to serialize output")
		res = workflow.SerializedExecutionResult("")
	}

	log.Debug().Msgf("Closing execution: %s", exId.String())
	err = s.execRepo.CloseExecution(ctx, exId, status, res)
	if err != nil {
		log.Error().Err(err).Msg("Failed to close workflow")
		return err
	}

	return nil
}

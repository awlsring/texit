package step_functions_workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *StepFunctionsWorkflow) DeprovisionNode(ctx context.Context, input *workflow.DeprovisionNodeInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Scheduling provision node workflow")

	err := s.scheduleWorkflow(ctx, s.deprovisionNodeWorkflowArn, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to schedule provision node workflow")
		return err
	}

	return nil
}

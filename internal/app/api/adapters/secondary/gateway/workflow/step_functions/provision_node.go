package step_functions_workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *StepFunctionsWorkflow) ProvisionNode(ctx context.Context, input *workflow.ProvisionNodeInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Scheduling provision node workflow")

	err := s.scheduleWorkflow(ctx, s.provisionNodeWorkflowArn, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to schedule provision node workflow")
		return err
	}

	return nil
}

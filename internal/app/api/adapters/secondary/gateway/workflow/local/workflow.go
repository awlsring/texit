package local_workflow

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type LocalWorkflow struct {
	workChan chan workflow.ExecutionInput
}

func New(c chan workflow.ExecutionInput) gateway.Workflow {
	return &LocalWorkflow{
		workChan: c,
	}
}

func (g *LocalWorkflow) DeprovisionNode(ctx context.Context, input *workflow.DeprovisionNodeInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Scheduling deprovision node workflow")

	err := g.scheduleWorkflow(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to schedule deprovision node workflow")
		return err
	}

	return nil
}

func (g *LocalWorkflow) ProvisionNode(ctx context.Context, input *workflow.ProvisionNodeInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Scheduling provision node workflow")

	err := g.scheduleWorkflow(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to schedule provision node workflow")
		return err
	}

	return nil
}

func (l *LocalWorkflow) scheduleWorkflow(ctx context.Context, input workflow.ExecutionInput) error {
	l.workChan <- input
	return nil
}

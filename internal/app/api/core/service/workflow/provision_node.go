package workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) LaunchProvisionNodeWorkflow(ctx context.Context, p *provider.Provider, l provider.Location, tn *tailnet.Tailnet, size node.Size, ephemeral bool) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Launching provision node workflow")

	exId := workflow.FormExecutionIdentifier(workflow.WorkflowNameProvisionNode)
	log.Debug().Msgf("Creating execution: %s", exId)
	ex := workflow.NewExecution(exId, workflow.WorkflowNameProvisionNode)

	input := &workflow.ProvisionNodeInput{
		ExecutionId:          exId.String(),
		ProviderName:         p.Name.String(),
		Location:             l.String(),
		TailnetName:          tn.Name.String(),
		TailnetControlServer: tn.ControlServer.String(),
		Size:                 size.String(),
		Ephemeral:            ephemeral,
	}

	log.Debug().Msg("Creating execution in repository")
	err := s.excRepo.CreateExecution(ctx, ex)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create execution")
		return "", err
	}

	log.Debug().Msg("Scheduling provision node workflow")
	err = s.wfGw.ProvisionNode(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to schedule provision node workflow")
		return "", err
	}

	return exId, nil
}

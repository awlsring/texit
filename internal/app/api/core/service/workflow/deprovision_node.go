package workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) LaunchDeprovisionNodeWorkflow(ctx context.Context, nid node.Identifier) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deprovisioning node")

	exId := workflow.FormExecutionIdentifier(workflow.WorkflowNameDeprovisionNode)
	log.Debug().Msgf("Creating execution: %s", exId)
	ex := workflow.NewExecution(exId, workflow.WorkflowNameDeprovisionNode)

	n, err := s.nodeRepo.Get(ctx, nid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return "", err
	}

	input := &workflow.DeprovisionNodeInput{
		ExecutionId:     exId.String(),
		NodeId:          n.Identifier.String(),
		Tailnet:         n.Tailnet.String(),
		TailnetDeviceId: n.TailnetIdentifier.String(),
	}

	log.Debug().Msg("Creating execution in repository")
	err = s.excRepo.CreateExecution(ctx, ex)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create execution")
		return "", err
	}

	log.Debug().Msg("Scheduling deprovision node workflow")
	err = s.wfGw.DeprovisionNode(ctx, input)
	if err != nil {
		log.Error().Err(err).Msg("Failed to schedule deprovision node workflow")
		return "", err
	}

	return exId, nil
}

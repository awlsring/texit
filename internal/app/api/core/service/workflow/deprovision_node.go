package workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) LaunchDeprovisionNodeWorkflow(ctx context.Context, id node.Identifier) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deprovisioning node")

	exId := workflow.FormExecutionIdentifier(workflow.WorkflowNameDeprovisionNode)

	log.Debug().Msgf("Creating execution: %s", exId)
	execution := workflow.NewExecution(exId, workflow.WorkflowNameDeprovisionNode)

	log.Debug().Msgf("Adding execution to database")
	err := s.excRepo.CreateExecution(ctx, execution)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create execution")
		return "", err
	}

	go func() {
		ctx = logger.InitContextLogger(context.Background(), log.GetLevel()) // TODO: Make workflow context logger

		log.Debug().Msg("Deleting node")
		log.Debug().Msgf("Getting node: %s", id)
		n, err := s.nodeRepo.Get(ctx, id)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get node")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to get node", err.Error()})
			return
		}

		log.Debug().Msgf("Getting platform gateway: %s", n.Provider)
		platformGw, err := s.getPlatformGateway(ctx, n.Provider)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get platform gateway")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to get platform gateway", err.Error()})
			return
		}

		log.Debug().Msg("Getting tailnet gateway")
		tailnetGw, err := s.getTailnetGateway(ctx, n.Tailnet)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get tailnet gateway")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to get tailnet gateway", err.Error()})
			return
		}

		log.Debug().Msgf("Deleting node from platform")
		err = platformGw.DeleteNode(ctx, n)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete node")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to delete node", err.Error()})
			return
		}

		log.Debug().Msgf("Deleting node from tailnet")
		err = tailnetGw.DeleteDevice(ctx, n.TailnetIdentifier)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to delete node from tailnet, continuing...")
		}

		log.Debug().Msg("Deleting node from repository")
		err = s.nodeRepo.Delete(ctx, id)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete node")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to delete node", err.Error()})
			return
		}

		log.Debug().Msg("Node deleted")
		s.closeWorkflow(ctx, exId, workflow.StatusComplete, nil)
	}()

	return exId, nil
}

package workflow

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (s *Service) LaunchDeprovisionNodeWorkflow(ctx context.Context, id node.Identifier) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deprovisioning node")

	exId := workflow.FormExecutionIdentifier(workflow.WorkflowNameDeprovisionNode)

	s.mu.Lock()
	execution := workflow.NewExecution(exId, workflow.WorkflowNameDeprovisionNode)
	s.executions[exId.String()] = execution
	s.mu.Unlock()

	go func() {
		ctx = logger.InitContextLogger(context.Background(), log.GetLevel()) // TODO: Make workflow context logger

		log.Debug().Msg("Deleting node")
		log.Debug().Msgf("Getting node: %s", id)
		n, err := s.nodeRepo.Get(ctx, id)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get node")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msgf("Getting platform gateway: %s", n.ProviderIdentifier)
		platformGw, err := s.getPlatformGateway(ctx, n.ProviderIdentifier)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get platform gateway")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msgf("Deleting node from platform: %s", platformGw)
		err = platformGw.DeleteNode(ctx, n)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete node")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msgf("Deleting node from tailnet")
		err = s.tailnetGw.DeleteDevice(ctx, n.TailnetIdentifier)
		if err != nil {
			log.Warn().Err(err).Msg("Failed to delete node from tailnet, continuing...")
		}

		log.Debug().Msg("Deleting node from repository")
		err = s.nodeRepo.Delete(ctx, id)
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete node")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msg("Node deleted")
		s.closeWorkflow(ctx, execution, workflow.StatusComplete)
	}()

	return exId, nil
}

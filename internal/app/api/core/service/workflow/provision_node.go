package workflow

import (
	"context"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

const (
	postCreationWaitTime = 5 * time.Second
)

func (s *Service) LaunchProvisionNodeWorkflow(ctx context.Context, provider provider.Identifier, location provider.Location, tn tailnet.Identifier, ephemeral bool) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Creating node on provider %s in location %s for tailnet %s", provider, location, tn)

	exId := workflow.FormExecutionIdentifier(workflow.WorkflowNameProvisionNode)

	s.mu.Lock()
	execution := workflow.NewExecution(exId, workflow.WorkflowNameProvisionNode)
	s.executions[exId.String()] = execution
	s.mu.Unlock()

	go func() {
		ctx = logger.InitContextLogger(context.Background(), log.GetLevel()) // TODO: Make workflow context logger

		log.Debug().Msg("Forming node id")
		id := node.FormNewNodeIdentifier()
		log.Debug().Msgf("New node id: %s", id)

		log.Debug().Msg("Forming tailnet identifier")
		tailName := tailnet.FormDeviceName(location.String(), id.String())
		log.Debug().Msgf("New tailnet device name: %s", tailName)

		log.Debug().Msg("Getting tailnet gateway")
		tailnetGw, err := s.getTailnetGateway(ctx, tn)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get tailnet gateway")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msg("Creating preauth key for node")
		preauthKey, err := tailnetGw.CreatePreauthKey(ctx, ephemeral)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create preauth key")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msg("Getting platfrom gateway")
		platformGw, err := s.getPlatformGateway(ctx, provider)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get platform gateway")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msg("Creating node on platform")
		platId, err := platformGw.CreateNode(ctx, id, tailName, provider, location, preauthKey)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create node")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msgf("Sleeping for %s to wait for device registration on tailnet", postCreationWaitTime)
		time.Sleep(postCreationWaitTime)

		log.Debug().Msg("Getting the tailnet device id")
		tid, err := tailnetGw.GetDeviceId(ctx, tailName)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get tailnet device id")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msg("Enabling as exit node")
		err = tailnetGw.EnableExitNode(ctx, tid)
		if err != nil {
			log.Error().Err(err).Msg("Failed to enable exit node")
		}

		log.Debug().Msg("Froming node entry")
		n := &node.Node{
			Identifier:         id,
			PlatformIdentifier: platId,
			Provider:           provider,
			Location:           location,
			PreauthKey:         preauthKey,
			Tailnet:            tn,
			TailnetIdentifier:  tid,
			Ephemeral:          ephemeral,
			TailnetName:        tailName,
		}

		log.Debug().Msg("Creating node in repository")
		err = s.nodeRepo.Create(ctx, n)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create node")
			s.closeWorkflow(ctx, execution, workflow.StatusFailed)
			return
		}

		log.Debug().Msgf("Node created, id: %s", id)
		s.closeWorkflow(ctx, execution, workflow.StatusComplete)
	}()

	return exId, nil
}

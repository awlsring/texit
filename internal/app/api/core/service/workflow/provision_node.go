package workflow

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

const (
	postCreationWaitTime = 5 * time.Second
)

func (s *Service) LaunchProvisionNodeWorkflow(ctx context.Context, prov *provider.Provider, location provider.Location, tn *tailnet.Tailnet, ephemeral bool) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Creating node on provider %s in location %s for tailnet %s", prov.Name.String(), location, tn)

	exId := workflow.FormExecutionIdentifier(workflow.WorkflowNameProvisionNode)

	log.Debug().Msgf("Creating execution: %s", exId)
	execution := workflow.NewExecution(exId, workflow.WorkflowNameProvisionNode)

	log.Debug().Msgf("Adding execution to database")
	err := s.excRepo.CreateExecution(ctx, execution)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create execution")
		return "", err
	}

	go func() {
		ctx = logger.InitContextLogger(context.Background(), log.GetLevel()) // TODO: Make workflow context logger

		log.Debug().Msg("Forming node id")
		id := node.FormNewNodeIdentifier()
		log.Debug().Msgf("New node id: %s", id)

		log.Debug().Msg("Forming tailnet identifier")
		tailName := tailnet.FormDeviceName(location.String(), id.String())
		log.Debug().Msgf("New tailnet device name: %s", tailName)

		log.Debug().Msg("Getting tailnet gateway")
		tailnetGw, err := s.getTailnetGateway(ctx, tn.Name)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get tailnet gateway")
			s.closeWorkflow(ctx, execution.Identifier, workflow.StatusFailed, []string{"Failed to get tailnet gateway", err.Error()})
			return
		}

		log.Debug().Msg("Creating preauth key for node")
		preauthKey, err := tailnetGw.CreatePreauthKey(ctx, ephemeral)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create preauth key")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to create preauth key", err.Error()})
			return
		}

		log.Debug().Msg("Getting platfrom gateway")
		platformGw, err := s.getPlatformGateway(ctx, prov.Name)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get platform gateway")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to get platform gateway", err.Error()})
			return
		}

		log.Debug().Msg("Creating node on platform")
		platId, err := platformGw.CreateNode(ctx, id, tailName, prov, location, tn, preauthKey)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create node")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to create node", err.Error()})
			return
		}

		log.Debug().Msgf("Sleeping for %s to wait for device registration on tailnet", postCreationWaitTime)
		time.Sleep(postCreationWaitTime)

		log.Debug().Msg("Getting the tailnet device id")
		tid, err := tailnetGw.GetDeviceId(ctx, tailName)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get tailnet device id")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to get tailnet device id", err.Error()})
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
			Provider:           prov.Name,
			Location:           location,
			PreauthKey:         preauthKey,
			Tailnet:            tn.Name,
			TailnetIdentifier:  tid,
			Ephemeral:          ephemeral,
			TailnetName:        tailName,
		}

		log.Debug().Msg("Creating node in repository")
		err = s.nodeRepo.Create(ctx, n)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create node")
			s.closeWorkflow(ctx, exId, workflow.StatusFailed, []string{"Failed to create node", err.Error()})
			return
		}

		log.Debug().Msgf("Node created, id: %s", id)
		s.closeWorkflow(ctx, exId, workflow.StatusComplete, nil)
	}()

	return exId, nil
}

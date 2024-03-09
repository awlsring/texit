package mem_workflow

import (
	"context"
	"errors"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

const (
	postCreationPollAmount = 40
	postCreationInterval   = 10 * time.Second
)

func returnFailure(err error, res workflow.ProvisionNodeExecutionResult) (workflow.Status, workflow.ExecutionResult) {
	res.SetError(err.Error())
	return workflow.StatusFailed, res
}

func (w *Worker) provisionNodeWorkflow(ctx context.Context, input *workflow.ProvisionNodeInput) (workflow.Status, workflow.ExecutionResult) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Provisioning node")

	results := workflow.NewProvisionNodeExecutionResult()

	log.Debug().Msg("Validating input")
	provName, err := provider.IdentifierFromString(input.ProviderName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		return returnFailure(err, results)
	}

	location, err := provider.LocationFromString(input.Location)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse location")
		return returnFailure(err, results)
	}

	tn, err := tailnet.IdentifierFromString(input.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet")
		return returnFailure(err, results)
	}

	tcs, err := tailnet.ControlServerFromString(input.TailnetControlServer)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet control server")
		return returnFailure(err, results)
	}

	size, err := node.SizeFromString(input.Size)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node size")
		return returnFailure(err, results)
	}

	log.Debug().Msg("Forming node id")
	id := node.FormNewNodeIdentifier()
	log.Debug().Msgf("New node id: %s", id)

	log.Debug().Msg("Forming tailnet identifier")
	tailName := tailnet.FormDeviceName(input.Location, id.String())
	log.Debug().Msgf("New tailnet device name: %s", tailName)

	log.Debug().Msg("Creating preauth key for node")
	preauthKey, err := w.actSvc.CreatePreauthKey(ctx, tn, input.Ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create preauth key")
		return returnFailure(err, results)
	}

	log.Debug().Msg("Creating node on platform")
	platId, err := w.actSvc.CreateNode(ctx, provName, tcs, id, tailName, location, preauthKey, size)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node")
		return returnFailure(err, results)
	}

	log.Debug().Msg("Getting the tailnet device id")
	var tid tailnet.DeviceIdentifier
	for i := 0; i < postCreationPollAmount; i++ {
		tid, err = w.actSvc.GetDeviceId(ctx, tn, tailName)
		if err != nil {
			if errors.Is(err, service.ErrUnknownTailnetDevice) {
				if i < postCreationPollAmount-1 {
					log.Debug().Msg("Device not found, sleeping and retrying")
					time.Sleep(postCreationInterval)
					continue
				} else {
					log.Error().Err(err).Msg("Failed to get tailnet device id")
					return returnFailure(err, results)
				}
			}
			log.Error().Err(err).Msg("Failed to get tailnet device id")
			return returnFailure(err, results)
		}
	}

	log.Debug().Msg("Enabling as exit node")
	err = w.actSvc.EnableExitNode(ctx, tn, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to enable exit node")
		return returnFailure(err, results)
	}

	log.Debug().Msg("Creating node in repository")
	err = w.actSvc.CreateNodeRecord(ctx, id, platId, provName, location, preauthKey, tn, tid, tailName, size, input.Ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node record")
		return returnFailure(err, results)
	}

	nodeId := id.String()
	results.Node = &nodeId
	log.Debug().Msgf("Node created, id: %s", id)
	return workflow.StatusComplete, results
}

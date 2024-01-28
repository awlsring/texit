package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) DeprovisionNode(ctx context.Context, req *teen.DeprovisionNodeRequest) (*teen.DeprovisionNodeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved delete node request")

	log.Debug().Msg("Validating delete node request")
	nodeId, err := node.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	log.Debug().Msg("Launching deprovision node workflow")
	exId, err := h.workSvc.LaunchDeprovisionNodeWorkflow(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to launch deprovision node workflow")
		return nil, err
	}

	log.Debug().Msg("Successfully started deprovision node")
	return &teen.DeprovisionNodeResponse{
		ExecutionId: exId.String(),
	}, nil
}

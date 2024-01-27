package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) StartNode(ctx context.Context, req *teen.StartNodeRequest) (*teen.StartNodeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved start node request")

	nodeId, err := node.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	err = h.nodeSvc.Start(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start node")
		return nil, err
	}

	log.Debug().Msg("Successfully started node")
	return &teen.StartNodeResponse{
		Success: true,
	}, nil
}

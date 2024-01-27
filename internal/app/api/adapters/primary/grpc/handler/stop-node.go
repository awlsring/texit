package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) StopNode(ctx context.Context, req *teen.StopNodeRequest) (*teen.StopNodeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved stop node request")

	nodeId, err := node.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	err = h.nodeSvc.Stop(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to stop node")
		return nil, err
	}

	log.Debug().Msg("Successfully started node")
	return &teen.StopNodeResponse{
		Success: true,
	}, nil
}

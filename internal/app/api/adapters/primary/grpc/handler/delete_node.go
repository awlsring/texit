package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) DeleteNode(ctx context.Context, req *teen.DeleteNodeRequest) (*teen.DeleteNodeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved delete node request")

	nodeId, err := node.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	err = h.nodeSvc.Delete(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node")
		return nil, err
	}

	log.Debug().Msg("Successfully deleted node")
	return &teen.DeleteNodeResponse{
		Success: true,
	}, nil
}

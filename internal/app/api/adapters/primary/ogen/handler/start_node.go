package handler

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) StartNode(ctx context.Context, req texit.StartNodeParams) (texit.StartNodeRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved start node request")

	nodeId, err := node.IdentifierFromString(req.Identifier)
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
	return &texit.StartNodeResponseContent{
		Success: true,
	}, nil
}

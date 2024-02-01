package handler

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) StopNode(ctx context.Context, req texit.StopNodeParams) (texit.StopNodeRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved stop node request")

	nodeId, err := node.IdentifierFromString(req.Identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	err = h.nodeSvc.Stop(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to stop node")
		return nil, err
	}

	log.Debug().Msg("Successfully stopped node")
	return &texit.StopNodeResponseContent{
		Success: true,
	}, nil
}

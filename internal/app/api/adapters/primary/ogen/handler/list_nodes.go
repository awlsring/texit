package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) ListNodes(ctx context.Context) (texit.ListNodesRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list nodes request")

	log.Debug().Msg("Listing nodes")
	nodes, err := h.nodeSvc.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes")
		return nil, err
	}

	log.Debug().Msg("Converting nodes to summaries")
	summaries := make([]texit.NodeSummary, len(nodes))
	for i, node := range nodes {
		summaries[i] = conversion.NodeToSummary(node)
	}

	log.Debug().Msg("Successfully listed nodes")
	return &texit.ListNodesResponseContent{
		Summaries: summaries,
	}, nil
}

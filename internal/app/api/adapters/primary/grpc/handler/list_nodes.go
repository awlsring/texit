package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/texit/internal/pkg/logger"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (h *Handler) ListNodes(ctx context.Context, _ *teen.ListNodesRequest) (*teen.ListNodesResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list nodes request")

	log.Debug().Msg("Listing nodes")
	nodes, err := h.nodeSvc.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes")
		return nil, err
	}

	log.Debug().Msg("Converting nodes to summaries")
	summaries := make([]*teen.NodeSummary, len(nodes))
	for i, node := range nodes {
		summaries[i] = conversion.NodeToSummary(node)
	}

	log.Debug().Msg("Successfully listed nodes")
	return &teen.ListNodesResponse{
		Nodes: summaries,
	}, nil
}

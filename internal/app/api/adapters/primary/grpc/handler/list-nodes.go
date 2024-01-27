package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) ListNodes(ctx context.Context, req *teen.ListNodesRequest) (*teen.ListNodesResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list nodes request")

	nodes, err := h.nodeSvc.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes")
		return nil, err
	}

	summaries := make([]*teen.NodeSummary, len(nodes))
	for i, node := range nodes {
		summaries[i] = conversion.NodeToSummary(node)
	}

	log.Debug().Msg("Successfully listed nodes")
	return &teen.ListNodesResponse{
		Nodes: summaries,
	}, nil
}

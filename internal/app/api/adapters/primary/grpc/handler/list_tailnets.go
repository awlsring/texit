package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/texit/internal/pkg/logger"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (h *Handler) ListTailnets(ctx context.Context, _ *teen.ListTailnetsRequest) (*teen.ListTailnetsResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list tailnets request")

	log.Debug().Msg("Listing tailnets")
	tns, err := h.tailnetSvc.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list tailnets")
		return nil, err
	}
	log.Debug().Msgf("Found %d tailnets", len(tns))
	if len(tns) == 0 {
		log.Warn().Msg("No tailnet found")
		return &teen.ListTailnetsResponse{}, nil
	}

	log.Debug().Msg("Converting tailnets to summaries")
	summaries := make([]*teen.TailnetSummary, len(tns))
	for i, t := range tns {
		summaries[i] = conversion.TailnetToSummary(t)
	}

	log.Debug().Msg("Successfully listed providers")
	return &teen.ListTailnetsResponse{
		Tailnets: summaries,
	}, nil
}

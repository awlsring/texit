package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) ListTailnets(ctx context.Context) (*texit.ListTailnetsResponseContent, error) {
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
		return &texit.ListTailnetsResponseContent{}, nil
	}

	log.Debug().Msg("Converting tailnets to summaries")
	summaries := make([]texit.TailnetSummary, len(tns))
	for i, t := range tns {
		summaries[i] = conversion.TailnetToSummary(t)
	}

	log.Debug().Msg("Successfully listed providers")
	return &texit.ListTailnetsResponseContent{
		Summaries: summaries,
	}, nil
}

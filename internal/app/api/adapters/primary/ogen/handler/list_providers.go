package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) ListProviders(ctx context.Context) (*texit.ListProvidersResponseContent, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list providers request")

	log.Debug().Msg("Listing providers")
	providers, err := h.providerSvc.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list providers")
		return nil, err
	}
	log.Debug().Msgf("Found %d providers", len(providers))
	if len(providers) == 0 {
		log.Warn().Msg("No providers found")
		return &texit.ListProvidersResponseContent{}, nil
	}

	log.Debug().Msg("Converting providers to summaries")
	summaries := make([]texit.ProviderSummary, len(providers))
	for i, provider := range providers {
		summaries[i] = conversion.ProviderToSummary(provider)
	}

	log.Debug().Msg("Successfully listed providers")
	return &texit.ListProvidersResponseContent{
		Summaries: summaries,
	}, nil
}

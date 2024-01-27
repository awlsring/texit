package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) ListProviders(ctx context.Context, _ *emptypb.Empty) (*teen.ListProvidersResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved list providers request")

	log.Debug().Msg("Listing providers")
	providers, err := h.providerSvc.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list providers")
		return nil, err
	}
	log.Debug().Msg("Successfully listed providers")

	log.Debug().Msg("Converting providers to summaries")
	summaries := make([]*teen.ProviderSummary, len(providers))
	for i, provider := range providers {
		summaries[i] = conversion.ProviderToSummary(provider)
	}

	log.Debug().Msg("Successfully listed providers")
	return &teen.ListProvidersResponse{
		Providers: summaries,
	}, nil
}

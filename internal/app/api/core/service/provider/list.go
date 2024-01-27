package provider

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (s *Service) List(ctx context.Context) ([]*provider.Provider, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("listing providers")

	providers := []*provider.Provider{}
	for _, p := range s.provMap {
		providers = append(providers, p)
	}

	log.Debug().Msgf("returning %d providers", len(providers))
	return providers, nil
}

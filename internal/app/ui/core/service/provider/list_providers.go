package provider

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) ListProviders(ctx context.Context) ([]*provider.Provider, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("listing providers")
	s.mut.Lock()
	if time.Since(s.lastRefresh) > defaultExpiration {
		log.Debug().Msg("refreshing providers")
		if err := s.refreshProviders(ctx); err != nil {
			log.Error().Err(err).Msg("failed to refresh providers")
			return nil, err
		}
	}
	s.mut.Unlock()
	log.Debug().Msg("listing providers from map")
	providers := make([]*provider.Provider, 0, len(s.providers))
	for _, p := range s.providers {
		providers = append(providers, p)
	}
	return providers, nil
}

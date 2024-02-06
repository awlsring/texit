package provider

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) DescribeProvider(ctx context.Context, id provider.Identifier) (*provider.Provider, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("describing provider")

	log.Debug().Msg("getting provider from map")
	tn, err := s.getProviderFromMap(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to get provider from map")
		return nil, err
	}

	log.Debug().Msg("returning provider")
	return tn, nil
}

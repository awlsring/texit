package tailnet

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) DescribeTailnet(ctx context.Context, id tailnet.Identifier) (*tailnet.Tailnet, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("describing tailnet")

	log.Debug().Msg("getting tailnet from map")
	tn, err := s.getTailnetFromMap(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to get tailnet from map")
		return nil, err
	}

	log.Debug().Msg("returning tailnet")
	return tn, nil
}

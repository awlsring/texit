package tailnet

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) ListTailnets(ctx context.Context) ([]*tailnet.Tailnet, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("listing tailnets")
	s.mut.Lock()
	if time.Since(s.lastRefresh) > defaultExpiration {
		log.Debug().Msg("refreshing tailnets")
		if err := s.refreshTailnets(ctx); err != nil {
			log.Error().Err(err).Msg("failed to refresh tailnets")
			return nil, err
		}
	}
	s.mut.Unlock()
	log.Debug().Msg("listing tailnets")
	tailnets := make([]*tailnet.Tailnet, 0, len(s.tailnets))
	for _, t := range s.tailnets {
		tailnets = append(tailnets, t)
	}
	return tailnets, nil
}

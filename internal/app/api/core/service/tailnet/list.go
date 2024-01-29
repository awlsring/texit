package tailnet

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (s *Service) List(ctx context.Context) ([]*tailnet.Tailnet, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("listing tailnets")

	tailnets := []*tailnet.Tailnet{}
	for _, t := range s.tailnetMap {
		tailnets = append(tailnets, t)
	}

	log.Debug().Msgf("returning %d providers", len(tailnets))
	return tailnets, nil
}

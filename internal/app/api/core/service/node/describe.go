package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Describe(ctx context.Context, id node.Identifier) (*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Describing node")

	log.Debug().Msgf("Getting node: %s", id)
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return nil, errors.Wrap(err, "failed to get node")
	}

	log.Debug().Msgf("Returning node: %s", n)
	return n, nil
}

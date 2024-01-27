package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) List(ctx context.Context) ([]*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Listing nodes")

	log.Debug().Msg("Getting nodes for repo")
	nodes, err := s.repo.List(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes")
		return nil, errors.Wrap(err, "failed to list nodes")
	}

	log.Debug().Msgf("Found %d nodes", len(nodes))
	return nodes, nil
}

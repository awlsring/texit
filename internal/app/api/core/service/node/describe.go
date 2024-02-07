package node

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Describe(ctx context.Context, id node.Identifier) (*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Describing node")

	log.Debug().Msgf("Getting node: %s", id)
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNodeNotFound) {
			log.Debug().Err(err).Msg("Node not found")
			return nil, service.ErrUnknownNode
		}
		log.Error().Err(err).Msg("Failed to get node")
		return nil, errors.Wrap(err, "failed to get node")
	}

	log.Debug().Msgf("Returning node: %s", n.Identifier)
	return n, nil
}

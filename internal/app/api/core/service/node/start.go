package node

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Start(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting node")

	log.Debug().Msg("Getting node")
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return errors.Wrap(err, "failed to get node")
	}

	log.Debug().Msg("Getting platform gateway")
	platformGw, err := s.getPlatformGateway(ctx, n.Provider)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return errors.Wrap(err, "failed to get platform gateway")
	}

	log.Debug().Msg("Starting node on platform")
	err = platformGw.StartNode(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start node")
		return errors.Wrap(err, "failed to start node")
	}

	log.Debug().Msg("Node started")
	return nil
}

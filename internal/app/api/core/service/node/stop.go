package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Stop(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Stopping node")

	log.Debug().Msg("Getting node")
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return errors.Wrap(err, "failed to get node")
	}

	log.Debug().Msg("Getting platform gateway")
	platformGw, err := s.getPlatformGateway(ctx, n.ProviderIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return errors.Wrap(err, "failed to get platform gateway")
	}

	log.Debug().Msg("Stopping node on platform")
	err = platformGw.StopNode(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to stop node")
		return errors.Wrap(err, "failed to stop node")
	}

	log.Debug().Msg("Node stopped")
	return nil
}

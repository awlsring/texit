package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Status(ctx context.Context, id node.Identifier) (node.Status, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting node status")

	log.Debug().Msgf("Getting node: %s", id)
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return node.StatusUnknown, errors.Wrap(err, "failed to get node")
	}

	log.Debug().Msgf("Getting platform gateway: %s", n.ProviderIdentifier)
	platform, err := s.getPlatformGateway(ctx, n.ProviderIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return node.StatusUnknown, errors.Wrap(err, "failed to get platform gateway")
	}

	log.Debug().Msgf("Getting node status from platform: %s", platform)
	status, err := platform.GetStatus(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node status")
		return node.StatusUnknown, errors.Wrap(err, "failed to get node status")
	}

	log.Debug().Msgf("Node status: %s", status)
	return status, nil
}

package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Delete(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting node")

	log.Debug().Msgf("Getting node: %s", id)
	n, err := s.repo.Get(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return errors.Wrap(err, "failed to get node")
	}

	log.Debug().Msgf("Getting platform gateway: %s", n.ProviderIdentifier)
	platformGw, err := s.getPlatformGateway(ctx, n.ProviderIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return errors.Wrap(err, "failed to get platform gateway")
	}

	log.Debug().Msgf("Deleting node from platform: %s", platformGw)
	err = platformGw.DeleteNode(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node")
		return errors.Wrap(err, "failed to delete node")
	}

	log.Debug().Msgf("Deleting node from repository: %s", s.repo)
	err = s.tailnetGw.DeleteDevice(ctx, n.TailnetIdentifier)
	if err != nil {
		return errors.Wrap(err, "failed to delete node")
	}

	log.Debug().Msgf("Deleting node from repository: %s", s.repo)
	err = s.repo.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node")
		return errors.Wrap(err, "failed to delete node")
	}

	log.Debug().Msg("Node deleted")
	return nil
}

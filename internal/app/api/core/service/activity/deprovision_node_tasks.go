package activity

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) DeleteNodeRecord(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Deleting node: %s", id)

	log.Debug().Msg("Deleting node record")
	err := s.nodeRepo.Delete(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node record")
		return err
	}

	return nil
}

func (s *Service) DeleteNode(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Deleting node: %s", id)

	log.Debug().Msg("Getting node")
	n, err := s.nodeRepo.Get(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node")
		return err
	}

	log.Debug().Msg("Getting platform gateway")
	platform, err := s.getPlatformGateway(ctx, n.Provider)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return err
	}

	log.Debug().Msg("Deleting node")
	err = platform.DeleteNode(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node")
		return err
	}

	return nil
}

func (s *Service) RemoveTailnetDevice(ctx context.Context, tailnet tailnet.Identifier, tailDeviceId tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Removing tailnet device: %s", tailnet)

	log.Debug().Msg("Getting tailnet gateway")
	gateway, err := s.getTailnetGateway(ctx, tailnet)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tailnet gateway")
		return err
	}

	log.Debug().Msg("Removing tailnet device")
	err = gateway.DeleteDevice(ctx, tailDeviceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to remove tailnet device")
		return err
	}

	return nil
}

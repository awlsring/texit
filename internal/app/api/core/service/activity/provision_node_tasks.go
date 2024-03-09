package activity

import (
	"context"
	"errors"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) CreateNodeRecord(ctx context.Context, nid node.Identifier, pid node.PlatformIdentifier, p provider.Identifier, l provider.Location, pk tailnet.PreauthKey, t tailnet.Identifier, tid tailnet.DeviceIdentifier, tn tailnet.DeviceName, si node.Size, e bool) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node record")

	now := time.Now()

	n := &node.Node{
		Identifier:         nid,
		PlatformIdentifier: pid,
		Provider:           p,
		Location:           l,
		PreauthKey:         pk,
		Tailnet:            t,
		TailnetIdentifier:  tid,
		TailnetName:        tn,
		Size:               si,
		Ephemeral:          e,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	log.Debug().Msg("Creating node record")
	err := s.nodeRepo.Create(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node record")
		return err
	}

	return nil
}

func (s *Service) CreateNode(ctx context.Context, provider provider.Identifier, tcs tailnet.ControlServer, node node.Identifier, tailName tailnet.DeviceName, location provider.Location, key tailnet.PreauthKey, size node.Size) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Creating node on provider %s in location %s", provider, location)

	log.Debug().Msg("Getting platform gateway")
	platform, err := s.getPlatformGateway(ctx, provider)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return "", err
	}

	log.Debug().Msg("Creating node")
	id, err := platform.CreateNode(ctx, node, tailName, location, tcs, key, size)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node")
		return "", err
	}

	return id, nil
}

func (s *Service) CreatePreauthKey(ctx context.Context, tailnet tailnet.Identifier, ephemeral bool) (tailnet.PreauthKey, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Creating preauth key for tailnet: %s", tailnet)

	log.Debug().Msg("Getting tailnet gateway")
	gateway, err := s.getTailnetGateway(ctx, tailnet)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tailnet gateway")
		return "", err
	}

	log.Debug().Msg("Creating preauth key")
	key, err := gateway.CreatePreauthKey(ctx, ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create preauth key")
		return "", err
	}

	return key, nil
}

func (s *Service) EnableExitNode(ctx context.Context, tailnet tailnet.Identifier, tailDeviceID tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Enabling exit node: %s", tailDeviceID)

	log.Debug().Msg("Getting tailnet gateway")
	gateway, err := s.getTailnetGateway(ctx, tailnet)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tailnet gateway")
		return err
	}

	log.Debug().Msg("Enabling exit node")
	err = gateway.EnableExitNode(ctx, tailDeviceID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to enable exit node")
		return err
	}

	return nil
}

func (s *Service) GetDeviceId(ctx context.Context, tailnet tailnet.Identifier, tailnetName tailnet.DeviceName) (tailnet.DeviceIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting device id for tailnet: %s, device: %s", tailnet, tailnetName)

	log.Debug().Msg("Getting tailnet gateway")
	gw, err := s.getTailnetGateway(ctx, tailnet)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tailnet gateway")
		return "", err
	}

	log.Debug().Msg("Getting device id")
	id, err := gw.GetDeviceId(ctx, tailnetName)
	if err != nil {
		if errors.Is(err, gateway.ErrUnknownDevice) {
			log.Warn().Err(err).Msg("Device not found")
			return "", service.ErrUnknownTailnetDevice
		}
		log.Error().Err(err).Msg("Failed to get device id")
		return "", err
	}

	return id, nil
}

package tailscale_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func (g *TailscaleGateway) EnableExitNode(ctx context.Context, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Info().Msgf("enabling exit node for %s", tid.String())

	log.Debug().Msg("updating acl")
	err := g.updateAcl(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to update acl")
		return err
	}

	log.Debug().Msg("setting device tags")
	err = g.client.SetDeviceTags(ctx, tid.String(), []string{tagTexitNode})
	if err != nil {
		log.Error().Err(err).Msg("failed to enable exit node")
		return err
	}

	log.Debug().Msg("disabling expiry")
	err = g.client.SetDeviceKey(ctx, tid.String(), tailscale.DeviceKey{
		KeyExpiryDisabled: true,
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to disable expiry")
	}

	log.Debug().Msg("exit node enabled")
	return nil
}

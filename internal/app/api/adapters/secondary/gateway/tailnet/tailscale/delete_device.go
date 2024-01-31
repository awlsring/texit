package tailscale_gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (g *TailscaleGateway) DeleteDevice(ctx context.Context, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Info().Msgf("Deleting device %s", tid.String())

	log.Debug().Msg("deleting device")
	err := g.client.DeleteDevice(ctx, tid.String())
	if err != nil {
		log.Error().Err(err).Msg("failed to delete device")
		return err
	}

	log.Debug().Msg("device deleted")
	return nil
}

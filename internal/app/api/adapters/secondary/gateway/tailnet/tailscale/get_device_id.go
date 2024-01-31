package tailscale_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (g *TailscaleGateway) GetDeviceId(ctx context.Context, tid tailnet.DeviceName) (tailnet.DeviceIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("getting device id for device with name %s", tid.String())

	log.Debug().Msg("listing devices")
	devices, err := g.client.Devices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to list devices")
		return "", err
	}

	log.Debug().Msg("searching for device")
	for _, device := range devices {
		if device.Hostname == tid.String() {
			log.Debug().Msgf("device found. id: %s", device.ID)
			deviceId, err := tailnet.DeviceIdentifierFromString(device.ID)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse device id")
				return "", err
			}
			return deviceId, nil
		}
	}

	log.Error().Msgf("device %s not found", tid.String())
	return "", errors.Wrap(gateway.ErrUnknownDevice, tid.String())
}

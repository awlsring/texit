package headscale_v0_22_3_gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/client/headscale_service"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

func (g *HeadscaleGateway) GetDeviceId(ctx context.Context, name tailnet.DeviceName) (tailnet.DeviceIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting headscale device id for device with name %s", name)

	log.Debug().Msg("Forming list devices request")
	request := headscale_service.NewHeadscaleServiceListMachinesParams()
	request.SetContext(ctx)
	request.SetUser(&g.user)

	log.Debug().Msg("Calling headscale")
	resp, err := g.client.HeadscaleServiceListMachines(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to get device id")
		return "", err
	}

	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("Finding device from list")
	for _, device := range resp.Payload.Machines {
		if device.Name == name.String() {
			log.Debug().Msgf("Device %s found", name.String())
			deviceId, err := tailnet.DeviceIdentifierFromString(device.ID)
			if err != nil {
				log.Error().Err(err).Msg("failed to parse device id")
				return "", err
			}
			return deviceId, nil
		}
	}

	log.Error().Msg("Device not found")
	return "", errors.Wrap(gateway.ErrUnknownDevice, name.String())
}

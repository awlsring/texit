package headscale_v0_22_3_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/client/headscale_service"
)

func (g *HeadscaleGateway) DeleteDevice(ctx context.Context, id tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Info().Msg("deleting device")

	log.Debug().Msg("creating headscale delete device request")
	request := headscale_service.NewHeadscaleServiceDeleteMachineParams()
	request.SetContext(ctx)
	request.SetMachineID(id.String())

	log.Debug().Msg("calling headscale")
	_, err := g.client.HeadscaleServiceDeleteMachine(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete device")
		return err
	}
	return nil
}

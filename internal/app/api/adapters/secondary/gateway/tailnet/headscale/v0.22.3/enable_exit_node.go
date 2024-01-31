package headscale_v0_22_3_gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/client/headscale_service"
)

func (g *HeadscaleGateway) getRoutesForDevice(ctx context.Context, tid tailnet.DeviceIdentifier) ([]string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("forming get routes request")
	request := headscale_service.NewHeadscaleServiceGetRoutesParams()
	request.SetContext(ctx)

	log.Debug().Msg("calling headscale")
	resp, err := g.client.HeadscaleServiceGetRoutes(request)
	if err != nil {
		log.Error().Err(err).Msg("failed to enable exit node")
		return nil, err
	}
	log.Debug().Msg("selecting routes exposed by device")
	targetRoutes := []string{}
	for _, route := range resp.Payload.Routes {
		if route.Machine.Name == tid.String() {
			log.Debug().Msgf("route %s exposed by device %s", route.ID, tid.String())
			targetRoutes = append(targetRoutes, route.ID)
		}
	}

	log.Debug().Msg("routes selected")
	return targetRoutes, nil
}

func (g *HeadscaleGateway) enableRoutes(ctx context.Context, routes []string) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("enabling routes")

	for _, route := range routes {
		log.Debug().Msgf("forming enable route request for route %s", route)
		request := headscale_service.NewHeadscaleServiceEnableRouteParams()
		request.SetContext(ctx)
		request.SetRouteID(route)

		log.Debug().Msg("calling headscale")
		_, err := g.client.HeadscaleServiceEnableRoute(request)
		if err != nil {
			log.Error().Err(err).Msg("failed to enable route")
			return err
		}
	}

	log.Debug().Msg("routes enabled")
	return nil
}

func (g *HeadscaleGateway) EnableExitNode(ctx context.Context, id tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Info().Msg("enabling exit node")

	log.Debug().Msg("forming get routes request")
	routes, err := g.getRoutesForDevice(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to enable exit node")
		return err
	}

	log.Debug().Msg("enabling routes")
	err = g.enableRoutes(ctx, routes)
	if err != nil {
		log.Error().Err(err).Msg("failed to enable exit node")
		return err
	}

	log.Debug().Msg("exit node enabled")
	return nil
}

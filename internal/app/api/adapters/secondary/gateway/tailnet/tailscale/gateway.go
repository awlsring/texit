package tailscale_gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type TailscaleGateway struct {
	client *tailscale.Client
}

func New(client *tailscale.Client) gateway.Tailnet {
	return &TailscaleGateway{
		client: client,
	}
}

func (g *TailscaleGateway) CreatePreauthKey(ctx context.Context) (tailnet.PreauthKey, error) {
	log := logger.FromContext(ctx)

	caps := tailscale.KeyCapabilities{
		Devices: struct {
			Create struct {
				Reusable      bool     `json:"reusable"`
				Ephemeral     bool     `json:"ephemeral"`
				Tags          []string `json:"tags"`
				Preauthorized bool     `json:"preauthorized"`
			} `json:"create"`
		}{
			Create: struct {
				Reusable      bool     `json:"reusable"`
				Ephemeral     bool     `json:"ephemeral"`
				Tags          []string `json:"tags"`
				Preauthorized bool     `json:"preauthorized"`
			}{
				Reusable:      false,
				Ephemeral:     false,
				Preauthorized: true,
			},
		},
	}

	log.Info().Msg("creating preauth key")
	resp, err := g.client.CreateKey(ctx, caps)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("preauth key created")
	return tailnet.PreauthKey(resp.Key), nil
}

func (g *TailscaleGateway) DeletePreauthKey(ctx context.Context, key tailnet.PreauthKey) error {
	log := logger.FromContext(ctx)

	log.Info().Msg("deleting preauth key")
	err := g.client.DeleteKey(ctx, key.String())
	if err != nil {
		return err
	}

	log.Debug().Msg("preauth key deleted")
	return nil
}

func (g *TailscaleGateway) DeleteDevice(ctx context.Context, id tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)

	log.Info().Msg("deleting device")
	err := g.client.DeleteDevice(ctx, id.String())
	if err != nil {
		return err
	}

	log.Debug().Msg("device deleted")
	return nil
}

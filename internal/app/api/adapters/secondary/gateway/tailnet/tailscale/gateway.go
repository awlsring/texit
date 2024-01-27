package tailscale

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

type TailscaleGateway struct {
	client *tailscale.Client
}

func New(client *tailscale.Client) *TailscaleGateway {
	return &TailscaleGateway{
		client: client,
	}
}

func (g *TailscaleGateway) CreatePreauthKey(ctx context.Context) (string, error) {
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
	return resp.Key, nil
}

func (g *TailscaleGateway) DeletePreauthKey(ctx context.Context, key string) error {
	log := logger.FromContext(ctx)

	log.Info().Msg("deleting preauth key")
	err := g.client.DeleteKey(ctx, key)
	if err != nil {
		return err
	}

	log.Debug().Msg("preauth key deleted")
	return nil
}

func (g *TailscaleGateway) DeleteDevice(ctx context.Context, id string) error {
	log := logger.FromContext(ctx)

	log.Info().Msg("deleting device")
	err := g.client.DeleteDevice(ctx, id)
	if err != nil {
		return err
	}

	log.Debug().Msg("device deleted")
	return nil
}

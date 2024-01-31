package tailscale_gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func (g *TailscaleGateway) CreatePreauthKey(ctx context.Context, ephemeral bool) (tailnet.PreauthKey, error) {
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
				Ephemeral:     ephemeral,
				Preauthorized: true,
			},
		},
	}

	log.Info().Msg("creating preauth key")
	resp, err := g.client.CreateKey(ctx, caps)
	if err != nil {
		log.Error().Err(err).Msg("failed to create preauth key")
		return "", err
	}

	log.Debug().Msg("preauth key created")
	return tailnet.PreauthKey(resp.Key), nil
}

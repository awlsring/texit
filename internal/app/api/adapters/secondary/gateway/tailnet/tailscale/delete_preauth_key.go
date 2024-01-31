package tailscale_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (g *TailscaleGateway) DeletePreauthKey(ctx context.Context, key tailnet.PreauthKey) error {
	log := logger.FromContext(ctx)

	log.Info().Msg("deleting preauth key")
	err := g.client.DeleteKey(ctx, key.String())
	if err != nil {
		log.Error().Err(err).Msg("failed to delete preauth key")
		return err
	}

	log.Debug().Msg("preauth key deleted")
	return nil
}

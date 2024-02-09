package platform_linode

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (p *PlatformLinode) StopNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Stopping Linode node")

	log.Debug().Msgf("Converting platform id %s to int", n.PlatformIdentifier)
	id, err := convertPlatformId(n.PlatformIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert Linode ID to int")
		return err
	}

	log.Debug().Msgf("Stopping Linode instance %s", n.PlatformIdentifier)
	err = p.client.ShutdownInstance(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to stop Linode instance")
		return err
	}
	log.Debug().Msg("Linode node stopped")
	return nil
}

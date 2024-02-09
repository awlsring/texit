package platform_linode

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (p *PlatformLinode) StartNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting Linode node")

	log.Debug().Msgf("Converting platform id %s to int", n.PlatformIdentifier)
	id, err := convertPlatformId(n.PlatformIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert Linode ID to int")
		return err
	}

	log.Debug().Msgf("Starting Linode instance %s", n.PlatformIdentifier)
	err = p.client.BootInstance(ctx, id, 0)
	if err != nil {
		log.Error().Err(err).Msg("Failed to start Linode instance")
		return err
	}
	log.Debug().Msg("Linode node started")
	return nil
}

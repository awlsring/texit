package platform_hetzner

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func (p *PlatformHetzner) DeleteNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting server")

	log.Debug().Msgf("Converting platform id %s to int64", n.PlatformIdentifier)
	id, err := convertPlatformId(n.PlatformIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert ID to int64")
		return err
	}

	log.Debug().Msgf("Deleting server %s", n.PlatformIdentifier)
	_, _, err = p.client.Server.DeleteWithResult(ctx, &hcloud.Server{ID: id})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete server")
		return err
	}
	log.Debug().Msg("server deleted")

	return nil
}

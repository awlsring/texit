package platform_hetzner

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func (p *PlatformHetzner) StartNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting hetzner node")

	log.Debug().Msgf("Converting platform id %s to int", n.PlatformIdentifier)
	id, err := convertPlatformId(n.PlatformIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert ID to int64")
		return err
	}

	log.Debug().Msgf("Starting server %s", n.PlatformIdentifier)
	_, _, err = p.client.Server.Poweron(ctx, &hcloud.Server{ID: id})
	if err != nil {
		log.Error().Err(err).Msg("Failed to start server")
		return err
	}
	log.Debug().Msg("Server started")
	return nil
}

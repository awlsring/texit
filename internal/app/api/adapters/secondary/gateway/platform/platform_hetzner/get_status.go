package platform_hetzner

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func translateState(status hcloud.ServerStatus) node.Status {
	switch status {
	case hcloud.ServerStatusDeleting, hcloud.ServerStatusStopping:
		return node.StatusStopping
	case hcloud.ServerStatusRunning:
		return node.StatusRunning
	case hcloud.ServerStatusInitializing, hcloud.ServerStatusStarting:
		return node.StatusStarting
	case hcloud.ServerStatusOff:
		return node.StatusStopped
	default:
		return node.StatusUnknown
	}

}

func (p *PlatformHetzner) GetStatus(ctx context.Context, n *node.Node) (node.Status, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting server status")

	log.Debug().Msgf("Getting server %s", n.PlatformIdentifier)
	s, _, err := p.client.Server.Get(ctx, n.PlatformIdentifier.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get server")
		return node.StatusUnknown, err
	}

	log.Debug().Msgf("server status is %s", s.Status)
	return translateState(s.Status), nil
}

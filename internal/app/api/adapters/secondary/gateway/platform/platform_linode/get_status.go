package platform_linode

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/linode/linodego"
)

func translateState(status linodego.InstanceStatus) node.Status {
	switch status {
	case linodego.InstanceOffline:
		return node.StatusStopped
	case linodego.InstanceRunning:
		return node.StatusRunning
	case linodego.InstanceRebooting:
		return node.StatusStarting
	case linodego.InstanceShuttingDown:
		return node.StatusStopping
	case linodego.InstanceProvisioning:
		return node.StatusStarting
	case linodego.InstanceBooting:
		return node.StatusStarting
	default:
		return node.StatusUnknown
	}

}

func (p *PlatformLinode) GetStatus(ctx context.Context, n *node.Node) (node.Status, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting Linode node status")

	log.Debug().Msgf("Converting platform id %s to int", n.PlatformIdentifier)
	id, err := convertPlatformId(n.PlatformIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert Linode ID to int")
		return node.StatusUnknown, err
	}

	log.Debug().Msgf("Getting Linode instance %s", n.PlatformIdentifier)
	instance, err := p.client.GetInstance(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get Linode instance")
		return node.StatusUnknown, err
	}

	log.Debug().Msgf("Linode node status is %s", instance.Status)
	return translateState(instance.Status), nil
}

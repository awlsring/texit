package platform_linode

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (p *PlatformLinode) getStackScriptId(ctx context.Context, name string) (int, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting stackscript id for %s", name)

	resp, err := p.client.ListStackscripts(ctx, nil)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list Linode stackscripts")
		return 0, err
	}

	for _, s := range resp {
		if s.Label == name {
			log.Debug().Msgf("Found stackscript id %d for %s", s.ID, name)
			return s.ID, nil
		}
	}

	log.Debug().Msgf("No stackscript found for %s", name)
	return 0, errors.New("no stackscript found")
}

func (p *PlatformLinode) DeleteNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting Linode node")

	log.Debug().Msgf("Converting platform id %s to int", n.PlatformIdentifier)
	id, err := convertPlatformId(n.PlatformIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert Linode ID to int")
		return err
	}

	log.Debug().Msgf("Deleting Linode instance %s", n.PlatformIdentifier)
	err = p.client.DeleteInstance(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete Linode instance")
		return err
	}
	log.Debug().Msg("Linode node deleted")

	log.Debug().Msgf("Getting stackscript id for %s", n.Identifier.String())
	stackId, err := p.getStackScriptId(ctx, n.Identifier.String())
	if err != nil {
		log.Warn().Err(err).Msg("Failed to get stackscript id")
		return nil
	}

	log.Debug().Msg("Deleting stackscript")
	err = p.client.DeleteStackscript(ctx, stackId)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to delete stackscript")
		return nil
	}
	log.Debug().Msg("Stackscript deleted")
	return nil
}

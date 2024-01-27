package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (s *Service) Create(ctx context.Context, provider provider.Identifier, location provider.Location) (*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node")

	log.Debug().Msg("Forming node id")
	id := node.FormNewNodeIdentifier()
	log.Debug().Msgf("New node id: %s", id)

	log.Debug().Msg("Forming tailnet identifier")
	tailId := tailnet.FormDeviceIdentifier(location.String(), id.String())

	log.Debug().Msg("Creating preauth key for node")
	preauthKey, err := s.tailnetGw.CreatePreauthKey(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create preauth key")
		return nil, err
	}

	log.Debug().Msg("Getting platfrom gateway")
	platformGw, err := s.getPlatformGateway(ctx, provider)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get platform gateway")
		return nil, err
	}

	log.Debug().Msg("Creating node on platform")
	platId, err := platformGw.CreateNode(ctx, id, tailId, provider, location, preauthKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node")
		return nil, err
	}

	log.Debug().Msg("Froming node entry")
	n := &node.Node{
		Identifier:         id,
		PlatformIdentifier: platId,
		ProviderIdentifier: provider,
		Location:           location,
		PreauthKey:         preauthKey,
		TailnetIdentifier:  tailId,
	}

	log.Debug().Msg("Creating node in repository")
	err = s.repo.Create(ctx, n)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node")
		return nil, err
	}

	log.Debug().Msg("Node created")
	return n, nil
}

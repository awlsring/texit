package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (g *PlatformAwsEcsGateway) CreateNode(ctx context.Context, _ node.Identifier, tid tailnet.DeviceIdentifier, pid provider.Identifier, loc provider.Location, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node on ECS")

	// TODO: Check if region is enabled
	log.Debug().Msgf("Getting ecs client for location %s", loc.String())
	ecsClient, err := g.getEcsClientForLocation(ctx, loc)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return "", err
	}

	log.Debug().Msgf("Getting default ec2 client for location %s", loc.String())
	ec2Client, err := g.getEc2ClientForLocation(ctx, loc)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return "", err
	}

	log.Debug().Msg("Creating task definition")
	err = g.createTaskDefinition(ctx, ecsClient, tid, key)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create task definition")
		return "", err
	}

	log.Debug().Msg("Making cluster")
	err = g.makeCluster(ctx, ecsClient)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create cluster")
		return "", err
	}

	log.Debug().Msg("Creating service")
	err = g.makeService(ctx, ecsClient, ec2Client, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create service")
		return "", err
	}

	log.Debug().Msg("Node created")
	return node.PlatformIdentifier(tid.String()), nil
}

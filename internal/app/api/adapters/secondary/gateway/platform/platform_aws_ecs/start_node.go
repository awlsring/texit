package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func (g *PlatformAwsEcsGateway) StartNode(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting ECS node")

	log.Debug().Msgf("Getting ecs client for location %s", node.Location.String())
	ecsClient, err := getClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, node.Location, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return err
	}

	log.Debug().Msg("Scaling service to 1 replica")
	err = scaleService(ctx, ecsClient, node.TailnetName, activeCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to scale service")
		return err
	}

	log.Debug().Msg("ECS node started")
	return nil
}

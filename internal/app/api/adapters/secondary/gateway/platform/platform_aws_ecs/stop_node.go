package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	platform_aws "github.com/awlsring/texit/internal/pkg/platform/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func (g *PlatformAwsEcsGateway) StopNode(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Stopping ECS node")

	log.Debug().Msgf("Getting ecs client for location %s", node.Location.String())
	ecsClient, err := platform_aws.GetClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, node.Location, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return err
	}

	log.Debug().Msg("Scaling service to 0 replicas")
	err = scaleService(ctx, ecsClient, node.Identifier, inactiveCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to scale service")
		return err
	}

	log.Debug().Msg("ECS node stopped")
	return nil
}

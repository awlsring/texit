package platform_aws_ecs

import (
	"context"

	platform_aws "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_common"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func (g *PlatformAwsEcsGateway) StartNode(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting ECS node")

	log.Debug().Msgf("Getting ecs client for location %s", node.Location.String())
	ecsClient, err := platform_aws.GetClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, node.Location, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return err
	}

	log.Debug().Msg("Scaling service to 1 replica")
	err = scaleService(ctx, ecsClient, node.Identifier, activeCount)
	if err != nil {
		log.Error().Err(err).Msg("Failed to scale service")
		return err
	}

	log.Debug().Msg("ECS node started")
	return nil
}

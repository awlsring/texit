package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// DeleteNode deletes all resources specific to the given node.
// If a resource doesn't generate a direct cost, if an error occurs it is logged and the process continues.
func (g *PlatformAwsEcsGateway) DeleteNode(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting ECS node")

	log.Debug().Msgf("Getting ecs client for location %s", node.Location.String())
	ecsClient, err := getClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, node.Location, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return err
	}

	log.Debug().Msgf("Getting ssm client for location %s", node.Location.String())
	ssmClient, err := getClientForLocation(ctx, ssm.NewFromConfig, g.ssmCache, node.Location, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return err
	}

	log.Debug().Msgf("Getting IAM client for location %s", node.Location.String())
	iamClient, err := getClientForLocation(ctx, iam.NewFromConfig, g.iamCache, node.Location, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get IAM client")
		return err
	}

	log.Debug().Msg("Deleting service")
	err = deleteService(ctx, ecsClient, node.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete service")
		return err
	}

	log.Debug().Msg("Deleting task definition")
	err = deleteTaskDefinition(ctx, ecsClient, node.TailnetName)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to delete task definition, continuing...")
	}

	log.Debug().Msg("Deleting state parameter")
	err = deleteStateParameter(ctx, ssmClient, node.TailnetName)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to delete state parameter, continuing...")
	}

	log.Debug().Msg("Deleting task role")
	err = deleteTaskRole(ctx, iamClient, node.TailnetName)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to delete task role, continuing...")
	}

	log.Debug().Msg("ECS node deleted")
	return nil
}

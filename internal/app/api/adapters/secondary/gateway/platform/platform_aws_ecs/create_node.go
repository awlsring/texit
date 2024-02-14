package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	platform_aws "github.com/awlsring/texit/internal/pkg/platform/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func (g *PlatformAwsEcsGateway) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, loc provider.Location, tcs tailnet.ControlServer, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node on ECS")

	// TODO: Check if region is enabled
	log.Debug().Msgf("Getting ecs client for location %s", loc.String())
	ecsClient, err := platform_aws.GetClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, loc, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return "", err
	}

	log.Debug().Msgf("Getting ec2 client for location %s", loc.String())
	ec2Client, err := platform_aws.GetClientForLocation(ctx, ec2.NewFromConfig, g.Ec2Cache, loc, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return "", err
	}

	log.Debug().Msgf("Getting iam client for location %s", loc.String())
	iamClient, err := platform_aws.GetClientForLocation(ctx, iam.NewFromConfig, g.iamCache, loc, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get IAM client")
		return "", err
	}

	log.Debug().Msgf("Getting SSM client for location %s", loc.String())
	ssmClient, err := platform_aws.GetClientForLocation(ctx, ssm.NewFromConfig, g.ssmCache, loc, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get IAM client")
		return "", err
	}

	log.Debug().Msg("Creating ECS execution role")
	execRole, err := makeExecutionRole(ctx, iamClient)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS execution role")
		return "", err
	}

	log.Debug().Msg("Creating SSM parameter")
	param, err := makeStateParameter(ctx, ssmClient, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create SSM parameter")
		return "", err
	}

	log.Debug().Msg("Creating ECS task role")
	taskRole, err := makeTaskRole(ctx, iamClient, id, param)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS task role")
		return "", err
	}

	log.Debug().Msg("Creating task definition")
	err = createTaskDefinition(ctx, ecsClient, id, tcs, tid, key, execRole, taskRole, param)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create task definition")
		return "", err
	}

	log.Debug().Msg("Making cluster")
	err = makeCluster(ctx, ecsClient)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create cluster")
		return "", err
	}

	log.Debug().Msg("Creating service")
	err = makeService(ctx, ecsClient, ec2Client, id, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create service")
		return "", err
	}

	log.Debug().Msg("Node created")
	return node.PlatformIdentifier(id.String()), nil
}

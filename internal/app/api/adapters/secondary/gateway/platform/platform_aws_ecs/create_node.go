package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func (g *PlatformAwsEcsGateway) CreateNode(ctx context.Context, _ node.Identifier, tid tailnet.DeviceName, pid *provider.Provider, loc provider.Location, tn *tailnet.Tailnet, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node on ECS")

	// TODO: Check if region is enabled
	log.Debug().Msgf("Getting ecs client for location %s", loc.String())
	ecsClient, err := getClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, loc, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return "", err
	}

	log.Debug().Msgf("Getting ec2 client for location %s", loc.String())
	ec2Client, err := getClientForLocation(ctx, ec2.NewFromConfig, g.ec2Cache, loc, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return "", err
	}

	log.Debug().Msgf("Getting iam client for location %s", loc.String())
	iamClient, err := getClientForLocation(ctx, iam.NewFromConfig, g.iamCache, loc, g.creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get IAM client")
		return "", err
	}

	log.Debug().Msgf("Getting SSM client for location %s", loc.String())
	ssmClient, err := getClientForLocation(ctx, ssm.NewFromConfig, g.ssmCache, loc, g.creds)
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
	param, err := makeStateParameter(ctx, ssmClient, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create SSM parameter")
		return "", err
	}

	log.Debug().Msg("Creating ECS task role")
	taskRole, err := makeTaskRole(ctx, iamClient, tid, param)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS task role")
		return "", err
	}

	log.Debug().Msg("Creating task definition")
	err = createTaskDefinition(ctx, ecsClient, tn, tid, key, execRole, taskRole, param)
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
	err = makeService(ctx, ecsClient, ec2Client, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create service")
		return "", err
	}

	log.Debug().Msg("Node created")
	return node.PlatformIdentifier(tid.String()), nil
}

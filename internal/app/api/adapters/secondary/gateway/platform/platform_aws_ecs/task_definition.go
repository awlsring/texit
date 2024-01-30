package platform_aws_ecs

import (
	"context"
	"fmt"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const (
	image                  = "ghcr.io/tailscale/tailscale:latest"
	defaultName            = "tailscale"
	defaultCpuAmount       = "256"
	defaultMemoryAmount    = "512"
	keyTsAuthkey           = "TS_AUTHKEY"
	keyTsStateDir          = "TS_STATE_DIR"
	keyTsHostname          = "TS_HOSTNAME"
	keyTsTailscaledArgs    = "TS_TAILSCALED_EXTRA_ARGS"
	keyTsExtraArgs         = "TS_EXTRA_ARGS"
	valueAdvertiseExitNode = "--advertise-exit-node"
	keyTsAcceptDns         = "TS_ACCEPT_DNS"
	keyTsUserspaceRoutes   = "TS_USERSPACE"
)

func createTaskDefinition(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceName, authkey tailnet.PreauthKey, execRole, taskRole, param string) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Creating new ECS task definition")
	_, err := client.RegisterTaskDefinition(ctx, &ecs.RegisterTaskDefinitionInput{
		ExecutionRoleArn: aws.String(execRole),
		TaskRoleArn:      aws.String(taskRole),
		ContainerDefinitions: []types.ContainerDefinition{
			{
				Name:      aws.String(defaultName),
				Image:     aws.String(image),
				Essential: aws.Bool(true),

				Environment: []types.KeyValuePair{
					{
						Name:  aws.String(keyTsAuthkey),
						Value: aws.String(authkey.String()),
					},
					{
						Name:  aws.String(keyTsHostname),
						Value: aws.String(tid.String()),
					},
					{
						Name:  aws.String(keyTsExtraArgs),
						Value: aws.String(valueAdvertiseExitNode),
					},
					{
						Name:  aws.String(keyTsAcceptDns),
						Value: aws.String("true"),
					},
					{
						Name:  aws.String(keyTsUserspaceRoutes),
						Value: aws.String("true"),
					},
					{
						Name:  aws.String(keyTsTailscaledArgs),
						Value: aws.String(fmt.Sprintf("--state=%s", param)),
					},
				},
			},
		},
		Tags: []types.Tag{
			{
				Key:   aws.String("created-by"),
				Value: aws.String("tailscale-cloud-exit-nodes"),
			},
			{
				Key:   aws.String("created-at"),
				Value: aws.String(time.Now().Format(time.RFC3339Nano)),
			},
			{
				Key:   aws.String("ephemeral"),
				Value: aws.String("true"),
			},
		},
		Family:      aws.String(tid.String()),
		Cpu:         aws.String(defaultCpuAmount),
		Memory:      aws.String(defaultMemoryAmount),
		NetworkMode: types.NetworkModeAwsvpc,
		RequiresCompatibilities: []types.Compatibility{
			types.CompatibilityFargate,
		},
		RuntimePlatform: &types.RuntimePlatform{
			CpuArchitecture:       types.CPUArchitectureArm64,
			OperatingSystemFamily: types.OSFamilyLinux,
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS task definition")
		return err
	}

	log.Debug().Msg("Created ECS task definition")
	return nil
}

func deleteTaskDefinition(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceName) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Deregistering ECS task definition")
	_, err := client.DeregisterTaskDefinition(ctx, &ecs.DeregisterTaskDefinitionInput{
		TaskDefinition: aws.String(taskDefinition(tid)),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete ECS task definition")
		return err
	}

	log.Debug().Msg("Deleting ECS task definition")
	_, err = client.DeleteTaskDefinitions(ctx, &ecs.DeleteTaskDefinitionsInput{
		TaskDefinitions: []string{fmt.Sprintf("%s:1", tid.String())},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete ECS task definition")
		return err
	}

	log.Debug().Msg("Deleted ECS task definition")
	return nil
}

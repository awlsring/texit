package platform_aws_ecs

import (
	"context"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/tailnet"
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
	keyTsHostname          = "TS_HOSTNAME"
	keyTsExtraArgs         = "TS_EXTRA_ARGS"
	valueAdvertiseExitNode = "--advertise-exit-node"
	keyTsAcceptDns         = "TS_ACCEPT_DNS"
	keyTsUserspaceRoutes   = "TS_USERSPACE"
)

func (g *PlatformAwsEcsGateway) createTaskDefinition(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier, authkey tailnet.PreauthKey) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Creating new ECS task definition")
	_, err := client.RegisterTaskDefinition(ctx, &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: []types.ContainerDefinition{
			{
				Name:      aws.String(defaultName),
				Image:     aws.String(image),
				Essential: aws.Bool(true),
				Hostname:  aws.String(tid.String()),
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
		},
		Family: aws.String(tid.String()),
		Cpu:    aws.String(defaultCpuAmount),
		Memory: aws.String(defaultMemoryAmount),
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

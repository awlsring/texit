package platform_aws_ecs

import (
	"context"
	"fmt"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const (
	activeCount            = 1
	inactiveCount          = 0
	servicePollMaxInterval = 10
	servicePollBackoff     = 10 * time.Second
	serviceStatusActive    = "ACTIVE"
)

func taskDefinition(tid tailnet.DeviceIdentifier) string {
	return fmt.Sprintf("%s:1", tid.String())
}

func (g *PlatformAwsEcsGateway) makeService(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Creating ECS service")
	_, err := client.CreateService(ctx, &ecs.CreateServiceInput{
		ServiceName:  aws.String(tid.String()),
		Cluster:      aws.String(clusterName),
		DesiredCount: aws.Int32(activeCount),
		LaunchType:   types.LaunchTypeFargate,
		CapacityProviderStrategy: []types.CapacityProviderStrategyItem{
			{
				CapacityProvider: aws.String(fargateCapacityProvider),
			},
		},
		TaskDefinition: aws.String(taskDefinition(tid)),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS service")
		return err
	}

	err = g.pollServiceTillCreated(ctx, client, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to poll ECS service")
		return err
	}

	log.Debug().Msg("ECS service created")
	return nil
}

func (g *PlatformAwsEcsGateway) pollServiceTillCreated(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Polling ECS service till it is created")
	for i := 0; i < servicePollMaxInterval; i++ {
		log.Debug().Msg("Polling ECS service")
		resp, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{
			Cluster:  aws.String(clusterName),
			Services: []string{tid.String()},
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to describe ECS service")
			return err
		}

		if len(resp.Services) == 0 {
			log.Error().Msg("Expected cluster not found")
			return err
		}

		if *resp.Services[0].Status == serviceStatusActive {
			log.Debug().Msg("ECS service is active")
			return nil
		}

		log.Debug().Msgf("ECS service not created, sleeping for %s", servicePollBackoff.String())
		time.Sleep(servicePollBackoff)
	}

	log.Error().Msg("ECS service not created")
	return nil
}

package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func scaleService(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier, desiredCount int32) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Scaling service")

	_, err := client.UpdateService(ctx, &ecs.UpdateServiceInput{
		Cluster:      aws.String(clusterName),
		Service:      aws.String(tid.String()),
		DesiredCount: aws.Int32(desiredCount),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to scale service")
		return err
	}

	log.Debug().Msg("Service scaled")
	return nil
}
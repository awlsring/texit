package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func deleteService(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting service")

	_, err := client.DeleteService(ctx, &ecs.DeleteServiceInput{
		Cluster: aws.String(clusterName),
		Service: aws.String(tid.String()),
		Force:   aws.Bool(true),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete service")
		return err
	}

	log.Debug().Msg("Service deleted")
	return nil
}

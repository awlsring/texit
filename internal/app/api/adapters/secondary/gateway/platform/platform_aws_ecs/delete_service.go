package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func deleteService(ctx context.Context, client interfaces.EcsClient, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting service")

	_, err := client.DeleteService(ctx, &ecs.DeleteServiceInput{
		Cluster: aws.String(clusterName),
		Service: aws.String(id.String()),
		Force:   aws.Bool(true),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete service")
		return err
	}

	log.Debug().Msg("Service deleted")
	return nil
}

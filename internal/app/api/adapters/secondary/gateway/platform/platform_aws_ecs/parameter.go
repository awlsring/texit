package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

func makeStateParameter(ctx context.Context, client interfaces.SsmClient, tid tailnet.DeviceIdentifier) (string, error) {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Creating state parameter")
	_, err := client.PutParameter(ctx, &ssm.PutParameterInput{
		Name:  aws.String(tid.String()),
		Type:  types.ParameterTypeString,
		Value: aws.String("{}"),
		Tags: []types.Tag{
			{
				Key:   aws.String("created-by"),
				Value: aws.String("tailscale-cloud-exit-nodes"),
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create state parameter")
		return "", err
	}
	log.Debug().Msg("State parameter created")

	log.Debug().Msg("Describe state parameter to get arn")
	resp, err := client.GetParameter(ctx, &ssm.GetParameterInput{
		Name: aws.String(tid.String()),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe state parameter")
		return "", err
	}

	return *resp.Parameter.ARN, nil
}

func deleteStateParameter(ctx context.Context, client interfaces.SsmClient, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting state parameter")

	_, err := client.DeleteParameter(ctx, &ssm.DeleteParameterInput{
		Name: aws.String(tid.String()),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete state parameter")
		return err
	}

	log.Debug().Msg("State parameter deleted")
	return nil
}

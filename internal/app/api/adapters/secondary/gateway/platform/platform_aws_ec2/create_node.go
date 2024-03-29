package platform_aws_ec2

import (
	"context"
	"encoding/base64"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/platform"
	platform_aws "github.com/awlsring/texit/internal/pkg/platform/aws"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func nodeSizeToInstanceTypeAndArch(size node.Size) (string, string) {
	switch size {
	case node.SizeSmall:
		return DefaultSmallInstanceType, DefaultSmallInstanceArch
	case node.SizeMedium:
		return DefaultMediumInstanceType, DefaultMediumInstanceArch
	case node.SizeLarge:
		return DefaultLargeInstanceType, DefaultLargeInstanceArch
	default:
		return DefaultSmallInstanceType, DefaultSmallInstanceArch
	}
}

func (g *PlatformAwsEc2Gateway) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, loc provider.Location, tcs tailnet.ControlServer, key tailnet.PreauthKey, size node.Size) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node on EC2")

	log.Debug().Msg("Getting instance type and architecture")
	instanceType, arch := nodeSizeToInstanceTypeAndArch(size)
	log.Debug().Msgf("Instance type: %s, architecture: %s", instanceType, arch)

	log.Debug().Msgf("Getting ec2 client for location %s", loc.String())
	ec2Client, err := platform_aws.GetClientForLocation(ctx, ec2.NewFromConfig, g.Ec2Cache, loc, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return "", err
	}

	log.Debug().Msgf("Getting latest AL2023 AMI for location %s", loc.String())
	ami, err := getLatestAmi(ctx, ec2Client, loc.String(), arch)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get latest AMI")
		return "", err
	}

	log.Debug().Msg("Forming cloud-init")
	cloudInit := formCloudInit(ctx, key.String(), tid, tcs)
	log.Debug().Str("cloud_init", cloudInit).Msg("Cloud-init formed")

	log.Debug().Msg("Creating EC2 instance")
	instanceId, err := runInstance(ctx, ec2Client, id, tid, ami, instanceType, cloudInit)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create EC2 instance")
		return "", err
	}
	log.Debug().Msgf("EC2 instance created with id %s", instanceId)

	log.Debug().Msg("Polling for instance to be ready")
	err = pollInstanceTillRunning(ctx, ec2Client, instanceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to poll for instance to be ready")
		return "", err
	}

	log.Debug().Msg("EC2 node created")
	return node.PlatformIdentifier(instanceId), nil
}

func runInstance(ctx context.Context, client interfaces.Ec2Client, id node.Identifier, tid tailnet.DeviceName, ami, typee, cloudInit string) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating EC2 instance")

	log.Debug().Msgf("Running instance with AMI %s and type %s", ami, typee)
	resp, err := client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      &ami,
		InstanceType: types.InstanceType(typee),
		UserData:     &cloudInit,
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
		ClientToken:  aws.String(id.String()),
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags: []types.Tag{
					{
						Key:   aws.String("created-by"),
						Value: aws.String("texit"),
					},
					{
						Key:   aws.String("node-id"),
						Value: aws.String(id.String()),
					},
					{
						Key:   aws.String("tailnet-device-name"),
						Value: aws.String(tid.String()),
					},
				},
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create EC2 instance")
		return "", err
	}

	log.Debug().Msg("EC2 instance created")
	return *resp.Instances[0].InstanceId, nil
}

func formCloudInit(ctx context.Context, authKey string, tid tailnet.DeviceName, tcs tailnet.ControlServer) string {
	userData := platform.TailscaleCloudInit(authKey, tid.String(), tcs.String())
	return base64.StdEncoding.EncodeToString([]byte(userData))
}

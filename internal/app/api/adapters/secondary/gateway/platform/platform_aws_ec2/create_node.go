package platform_aws_ec2

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	platform_aws "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_common"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const (
	postCreationSleep = 30
)

func (g *PlatformAwsEc2Gateway) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, pid *provider.Provider, loc provider.Location, tn *tailnet.Tailnet, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node on EC2")

	log.Debug().Msgf("Getting ec2 client for location %s", loc.String())
	ec2Client, err := platform_aws.GetClientForLocation(ctx, ec2.NewFromConfig, g.Ec2Cache, loc, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return "", err
	}

	log.Debug().Msgf("Getting latest AL2023 AMI for location %s", loc.String())
	ami, err := getLatestAmi(ctx, ec2Client, loc.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get latest AMI")
		return "", err
	}

	log.Debug().Msg("Forming cloud-init")
	cloudInit := formCloudInit(ctx, key.String(), loc.String(), tid)
	log.Debug().Str("cloud_init", cloudInit).Msg("Cloud-init formed")

	log.Debug().Msg("Creating EC2 instance")
	instanceId, err := runInstance(ctx, ec2Client, ami, DefaultInstanceType, cloudInit)
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

	log.Debug().Msg("EC2 node created, sleeping while it inits")
	time.Sleep(postCreationSleep * time.Second)

	log.Debug().Msg("EC2 node created")
	return node.PlatformIdentifier(instanceId), nil
}

func runInstance(ctx context.Context, client interfaces.Ec2Client, ami, typee, cloudInit string) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating EC2 instance")

	log.Debug().Msgf("Running instance with AMI %s and type %s", ami, typee)
	resp, err := client.RunInstances(ctx, &ec2.RunInstancesInput{
		ImageId:      &ami,
		InstanceType: types.InstanceType(typee),
		UserData:     &cloudInit,
		MaxCount:     aws.Int32(1),
		MinCount:     aws.Int32(1),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create EC2 instance")
		return "", err
	}

	log.Debug().Msg("EC2 instance created")
	return *resp.Instances[0].InstanceId, nil
}

func formCloudInit(ctx context.Context, authKey, location string, tid tailnet.DeviceName) string {
	cloudInit := fmt.Sprintf("#!/bin/bash\necho 'net.ipv4.ip_forward = 1' | sudo tee -a /etc/sysctl.conf\necho 'net.ipv6.conf.all.forwarding = 1' | sudo tee -a /etc/sysctl.conf\nsudo sysctl -p /etc/sysctl.conf\n\ncurl -fsSL https://tailscale.com/install.sh | sh\n\nsudo tailscale up --auth-key=%s --hostname=%s --advertise-exit-node", authKey, tid.String())
	return base64.StdEncoding.EncodeToString([]byte(cloudInit))
}

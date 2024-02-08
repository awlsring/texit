package platform_aws_ec2

import (
	"context"

	platform_aws "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_common"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func (g *PlatformAwsEc2Gateway) StopNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Stopping EC2 instance")

	log.Debug().Msgf("Getting ec2 client for location %s", n.Location.String())
	ec2Client, err := platform_aws.GetClientForLocation(ctx, ec2.NewFromConfig, g.Ec2Cache, n.Location, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return err
	}

	log.Debug().Msgf("Stopping EC2 instance %s", n.PlatformIdentifier)
	_, err = ec2Client.StopInstances(ctx, &ec2.StopInstancesInput{
		InstanceIds: []string{n.PlatformIdentifier.String()},
	})

	log.Debug().Msg("EC2 node stopped")
	return err
}

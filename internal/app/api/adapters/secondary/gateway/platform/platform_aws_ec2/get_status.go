package platform_aws_ec2

import (
	"context"

	platform_aws "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_common"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func translateEc2Status(status string) node.Status {
	switch status {
	case "running":
		return node.StatusRunning
	case "stopped":
		return node.StatusStopped
	case "stopping":
		return node.StatusStopping
	case "pending":
		return node.StatusStarting
	default:
		return node.StatusUnknown
	}
}

func (g *PlatformAwsEc2Gateway) GetStatus(ctx context.Context, n *node.Node) (node.Status, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting instance status")

	log.Debug().Msgf("Getting ec2 client for location %s", n.Location.String())
	ec2Client, err := platform_aws.GetClientForLocation(ctx, ec2.NewFromConfig, g.Ec2Cache, n.Location, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get EC2 client")
		return node.StatusUnknown, err
	}

	log.Debug().Msgf("Getting instance status for instance %s", n.PlatformIdentifier)
	status, err := getInstanceStatus(ctx, ec2Client, n.PlatformIdentifier.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to get instance status")
		return node.StatusUnknown, err
	}

	log.Debug().Msgf("Instance status: %s", status)
	return translateEc2Status(status), nil
}

package platform_aws_ec2

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const (
	defaultAmiFilter = "al2023-ami-*"
)

func getLatestAmi(ctx context.Context, client interfaces.Ec2Client, region string, arch string) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting latest AMI")

	amazonLinuxFilter := []types.Filter{
		{
			Name:   aws.String("name"),
			Values: []string{defaultAmiFilter},
		},
		{
			Name:   aws.String("architecture"),
			Values: []string{arch},
		},
	}

	resp, err := client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Filters: amazonLinuxFilter,
		Owners:  []string{"amazon"},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get latest AMI")
		return "", err
	}

	if len(resp.Images) == 0 {
		log.Debug().Msg("No AMI found")
		return "", errors.New("no AMI found")
	}

	latestAmi := resp.Images[0].ImageId
	log.Debug().Msgf("Latest AMI found: %s", *latestAmi)
	return *latestAmi, nil
}

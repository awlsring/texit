package platform_aws_ec2

import (
	"context"
	"errors"
	"time"

	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

const (
	defaultPollInterval = 50
	defaultPollBackoff  = 2 * time.Second
)

func getInstanceStatus(ctx context.Context, client interfaces.Ec2Client, id string) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting instance status")

	log.Debug().Msgf("Getting instance status for instance %s", id)
	resp, err := client.DescribeInstances(ctx, &ec2.DescribeInstancesInput{
		InstanceIds: []string{id},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get instance status")
		return "", err
	}

	status := string(resp.Reservations[0].Instances[0].State.Name)
	log.Debug().Msgf("Instance status: %s", status)
	return status, nil
}

func pollInstanceTillRunning(ctx context.Context, client interfaces.Ec2Client, id string) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Polling instance till running")

	for i := 0; i < defaultPollInterval; i++ {
		status, err := getInstanceStatus(ctx, client, id)
		if err != nil {
			return err
		}

		if status == "running" {
			log.Debug().Msg("Instance running")
			return nil
		}

		log.Debug().Msgf("Instance status: %s", status)
		time.Sleep(defaultPollBackoff)
	}

	log.Debug().Msg("Instance not running")
	return errors.New("instance not running")
}

package platform_aws_ecs

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	etypes "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const (
	activeCount            = 1
	inactiveCount          = 0
	servicePollMaxInterval = 40
	servicePollBackoff     = 2 * time.Second
	serviceStatusActive    = "ACTIVE"
	taskPollBackoff        = 2 * time.Second
	taskPollMaxInterval    = 40
	taskStatusActive       = "RUNNING"
)

func taskDefinition(tid tailnet.DeviceIdentifier) string {
	return fmt.Sprintf("%s:1", tid.String())
}

func getDefaultVpc(ctx context.Context, client interfaces.Ec2Client) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting default VPC")

	resp, err := client.DescribeVpcs(ctx, &ec2.DescribeVpcsInput{
		Filters: []etypes.Filter{
			{
				Name:   aws.String("isDefault"),
				Values: []string{"true"},
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe VPCs")
		return "", err
	}

	if len(resp.Vpcs) == 0 {
		log.Error().Msg("Expected default VPC not found")
		return "", err
	}

	log.Debug().Msg("Got default VPC")
	return *resp.Vpcs[0].VpcId, nil
}

func getDefaultSubnets(ctx context.Context, client interfaces.Ec2Client, vpc string) ([]string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting default subnets")

	resp, err := client.DescribeSubnets(ctx, &ec2.DescribeSubnetsInput{
		Filters: []etypes.Filter{
			{
				Name:   aws.String("vpc-id"),
				Values: []string{vpc},
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe subnets")
		return nil, err
	}

	if len(resp.Subnets) == 0 {
		log.Error().Msg("Expected default subnets not found")
		return nil, err
	}

	subnets := []string{}
	for _, subnet := range resp.Subnets {
		subnets = append(subnets, *subnet.SubnetId)
	}
	return subnets, nil
}

func getDefaultSecurityGroupId(ctx context.Context, client interfaces.Ec2Client) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting default security group")

	resp, err := client.DescribeSecurityGroups(ctx, &ec2.DescribeSecurityGroupsInput{
		Filters: []etypes.Filter{
			{
				Name:   aws.String("group-name"),
				Values: []string{"default"},
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe security groups")
		return "", err
	}

	if len(resp.SecurityGroups) == 0 {
		log.Error().Msg("Expected default security group not found")
		return "", err
	}

	log.Debug().Msg("Got default security group")
	return *resp.SecurityGroups[0].GroupId, nil
}

func makeService(ctx context.Context, ecsClient interfaces.EcsClient, ec2Client interfaces.Ec2Client, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Making ECS service")

	log.Debug().Msg("Getting default VPC")
	vpc, err := getDefaultVpc(ctx, ec2Client)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get default VPC")
		return err
	}

	log.Debug().Msg("Getting default subnets")
	subnets, err := getDefaultSubnets(ctx, ec2Client, vpc)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get default subnets")
		return err
	}

	log.Debug().Msg("Getting default security group")
	sgId, err := getDefaultSecurityGroupId(ctx, ec2Client)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get default security group")
		return err
	}

	log.Debug().Msg("Creating ECS service")
	_, err = ecsClient.CreateService(ctx, &ecs.CreateServiceInput{
		ServiceName:  aws.String(tid.String()),
		Cluster:      aws.String(clusterName),
		DesiredCount: aws.Int32(activeCount),
		LaunchType:   types.LaunchTypeFargate,
		NetworkConfiguration: &types.NetworkConfiguration{
			AwsvpcConfiguration: &types.AwsVpcConfiguration{
				AssignPublicIp: types.AssignPublicIpEnabled,
				Subnets:        subnets,
				SecurityGroups: []string{sgId},
			},
		},
		TaskDefinition: aws.String(taskDefinition(tid)),
		Tags: []types.Tag{
			{
				Key:   aws.String("created-by"),
				Value: aws.String("tailscale-cloud-exit-nodes"),
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS service")
		return err
	}

	err = pollServiceTillCreated(ctx, ecsClient, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to poll ECS service")
		return err
	}

	task, err := getLaunchedTask(ctx, ecsClient, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get launched ECS task")
		return err
	}

	err = pollTillTaskIsActive(ctx, ecsClient, task)
	if err != nil {
		log.Error().Err(err).Msg("Failed to poll ECS task")
		return err
	}

	log.Debug().Msg("ECS service created")
	return nil
}

func getLaunchedTask(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting launched ECS task")

	for i := 0; i < taskPollMaxInterval; i++ {
		resp, err := client.ListTasks(ctx, &ecs.ListTasksInput{
			Cluster:     aws.String(clusterName),
			ServiceName: aws.String(tid.String()),
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to list ECS tasks")
			return "", err
		}
		if len(resp.TaskArns) == 1 {
			log.Debug().Msg("Got launched ECS task")
			return resp.TaskArns[0], nil
		}

		log.Debug().Msgf("ECS task not launched, sleeping for %s", taskPollBackoff.String())
		time.Sleep(taskPollBackoff)
	}

	log.Error().Msg("ECS task not launched")
	return "", errors.New("ECS task not launched")
}

func pollTillTaskIsActive(ctx context.Context, client interfaces.EcsClient, task string) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Polling ECS task till it is active")
	for i := 0; i < taskPollMaxInterval; i++ {
		log.Debug().Msg("Polling ECS task")
		resp, err := client.DescribeTasks(ctx, &ecs.DescribeTasksInput{
			Cluster: aws.String(clusterName),
			Tasks:   []string{task},
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to describe ECS task")
			return err
		}

		if len(resp.Tasks) == 0 {
			log.Error().Msg("Expected task not found")
			return err
		}

		if *resp.Tasks[0].LastStatus == taskStatusActive {
			log.Debug().Msg("ECS task is active")
			return nil
		}

		log.Debug().Msgf("ECS task not active, sleeping for %s", taskPollBackoff.String())
		time.Sleep(taskPollBackoff)
	}

	log.Error().Msg("ECS task not active")
	return errors.New("ECS task not active within wait time")
}

func pollServiceTillCreated(ctx context.Context, client interfaces.EcsClient, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Polling ECS service till it is created")
	for i := 0; i < servicePollMaxInterval; i++ {
		log.Debug().Msg("Polling ECS service")
		resp, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{
			Cluster:  aws.String(clusterName),
			Services: []string{tid.String()},
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to describe ECS service")
			return err
		}

		if len(resp.Services) == 0 {
			log.Error().Msg("Expected cluster not found")
			return err
		}

		if *resp.Services[0].Status == serviceStatusActive {
			log.Debug().Msg("ECS service is active")
			return nil
		}

		log.Debug().Msgf("ECS service not created, sleeping for %s", servicePollBackoff.String())
		time.Sleep(servicePollBackoff)
	}

	log.Error().Msg("ECS service not created")
	return errors.New("ECS service not created")
}

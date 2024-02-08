package platform_aws_ecs

import (
	"context"
	"fmt"

	platform_aws "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_common"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func taskStatusToNodeStatus(status string) node.Status {
	switch status {
	case "RUNNING":
		return node.StatusRunning
	case "PENDING", "ACTIVATING":
		return node.StatusStarting
	case "STOPPED", "DELETED":
		return node.StatusStopped
	case "STOPPING", "DEPROVISIONING":
		return node.StatusStopping
	default:
		return node.StatusUnknown
	}
}

func listServiceTasks(ctx context.Context, client interfaces.EcsClient, id node.Identifier) ([]string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Listing service tasks")

	resp, err := client.ListTasks(ctx, &ecs.ListTasksInput{
		Cluster:     aws.String(clusterName),
		ServiceName: aws.String(id.String()),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to list service tasks")
		return nil, err
	}

	log.Debug().Msg("Service tasks listed")
	return resp.TaskArns, nil
}

func describeTaskStatus(ctx context.Context, client interfaces.EcsClient, arn string) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Describing task status")

	resp, err := client.DescribeTasks(ctx, &ecs.DescribeTasksInput{
		Cluster: aws.String(clusterName),
		Tasks:   []string{arn},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe task")
		return "", err
	}
	if len(resp.Tasks) == 0 {
		log.Debug().Msg("No tasks found")
		return "", fmt.Errorf("no tasks found")
	}
	log.Debug().Msg("Task status described")
	task := resp.Tasks[0]

	log.Debug().Msgf("Checking if task status is set")
	if task.LastStatus == nil {
		log.Debug().Msg("Task status not set")
		return "", nil
	}
	log.Debug().Msgf("Task status set, is %s", *task.LastStatus)
	return *task.LastStatus, nil
}

func (g *PlatformAwsEcsGateway) GetStatus(ctx context.Context, n *node.Node) (node.Status, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting service status")

	log.Debug().Msgf("Getting ecs client for location %s", n.Location.String())
	ecsClient, err := platform_aws.GetClientForLocation(ctx, ecs.NewFromConfig, g.ecsCache, n.Location, g.Creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ECS client")
		return node.StatusUnknown, err
	}

	tasks, err := listServiceTasks(ctx, ecsClient, n.Identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list service tasks")
		return node.StatusUnknown, err
	}
	if len(tasks) == 0 {
		log.Debug().Msg("No tasks found")
		return node.StatusStopped, nil
	}

	log.Debug().Msg("Getting task status")
	taskStatus, err := describeTaskStatus(ctx, ecsClient, tasks[0])
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe task status")
		return node.StatusUnknown, err
	}

	log.Debug().Msgf("Converting task status to node status")
	status := taskStatusToNodeStatus(taskStatus)

	log.Debug().Msgf("Service status is %s", status)
	return status, nil
}

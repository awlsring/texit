package platform_aws_ecs

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

const (
	fargateCapacityProvider = "FARGATE"
	clusterName             = "tailscale-cloud-exit-nodes"
	clusterActive           = "ACTIVE"
	clusterPollBackoff      = 10 * time.Second
	clusterPollMaxInterval  = 10
)

func (g *PlatformAwsEcsGateway) makeCluster(ctx context.Context, client interfaces.EcsClient) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Checking if ECS cluster exists")
	exists, err := g.checkIfClusterExists(ctx, client)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check if cluster exists")
		return err
	}
	if exists {
		log.Debug().Msg("Cluster exists")
		return nil
	}

	log.Debug().Msg("Cluster does not exist, creating")
	err = g.createNewCluster(ctx, client)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create cluster")
		return err
	}

	log.Debug().Msg("Polling cluster till it is created")
	err = g.pollClusterTillCreated(ctx, client)
	if err != nil {
		log.Error().Err(err).Msg("error while polling cluster")
		return err
	}

	log.Debug().Msg("Cluster created")
	return nil
}

func (g *PlatformAwsEcsGateway) createNewCluster(ctx context.Context, client interfaces.EcsClient) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Creating new ECS cluster")
	_, err := client.CreateCluster(ctx, &ecs.CreateClusterInput{
		CapacityProviders: []string{fargateCapacityProvider},
		ClusterName:       aws.String(clusterName),
		DefaultCapacityProviderStrategy: []types.CapacityProviderStrategyItem{
			{
				CapacityProvider: aws.String(fargateCapacityProvider),
			},
		},
		Tags: []types.Tag{
			{
				Key:   aws.String("created-by"),
				Value: aws.String("tailscale-cloud-exit-nodes"),
			},
			{
				Key:   aws.String("created-at"),
				Value: aws.String(time.Now().Format(time.RFC3339Nano)),
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (g *PlatformAwsEcsGateway) pollClusterTillCreated(ctx context.Context, client interfaces.EcsClient) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Polling ECS cluster till it is created")
	i := 0
	for {
		log.Debug().Msg("Checking if cluster exists")
		resp, err := client.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: []string{clusterName},
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to describe ECS clusters")
			return err
		}

		if len(resp.Clusters) == 0 {
			log.Error().Msg("Expected cluster not found")
			return err
		}

		if *resp.Clusters[0].Status == clusterActive {
			log.Debug().Msg("Cluster is active")
			return nil
		}

		log.Debug().Msg("Cluster does not exist, sleeping for 5 seconds")
		time.Sleep(clusterPollBackoff)
		i++
		if i > clusterPollMaxInterval {
			log.Error().Msg("Cluster did not become active in time")
			return errors.New("cluster did not become active in time")
		}
	}
}

func (g *PlatformAwsEcsGateway) checkIfClusterExists(ctx context.Context, client interfaces.EcsClient) (bool, error) {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Describing ECS clusters")
	resp, err := client.ListClusters(ctx, &ecs.ListClustersInput{})
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe ECS clusters")
		return false, err
	}

	log.Debug().Msgf("Clusters: %v", resp)
	for _, cluster := range resp.ClusterArns {
		if strings.Contains(cluster, clusterName) {
			return true, nil
		}
	}

	log.Debug().Msgf("Expected cluster with name %s not found", clusterName)
	return false, nil
}

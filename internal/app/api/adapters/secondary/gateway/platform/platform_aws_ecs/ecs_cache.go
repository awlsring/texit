package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/patrickmn/go-cache"
)

func (g *PlatformAwsEcsGateway) getEcsClientForLocation(ctx context.Context, loc provider.Location) (interfaces.EcsClient, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting ECS client for location %s", loc.String())

	var c interfaces.EcsClient
	if i, found := g.ecsCache.Get(loc.String()); found {
		log.Debug().Msg("Found ECS client in cache")
		c = i.(*ecs.Client)
		return c, nil
	}

	log.Debug().Msg("ECS client not found in cache, creating a new one")
	client, err := g.createEcsClientForLocation(ctx, loc)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS client")
		return nil, err
	}

	log.Debug().Msg("Returning ECS client")
	return client, nil
}

func (g *PlatformAwsEcsGateway) createEcsClientForLocation(ctx context.Context, loc provider.Location) (interfaces.EcsClient, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting ECS client for location %s", loc.String())

	log.Debug().Msg("Creating ECS client")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(loc.String()), config.WithCredentialsProvider(g.creds))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create ECS client")
		return nil, err
	}

	log.Debug().Msg("Creating ECS client")
	client := ecs.NewFromConfig(cfg)

	log.Debug().Msg("Caching ECS client")
	g.ecsCache.Set(loc.String(), &client, cache.DefaultExpiration)

	log.Debug().Msg("Returning ECS client")
	return client, nil
}

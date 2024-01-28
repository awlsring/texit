package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/patrickmn/go-cache"
)

func (g *PlatformAwsEcsGateway) getEc2ClientForLocation(ctx context.Context, loc provider.Location) (interfaces.Ec2Client, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting EC2 client for location %s", loc.String())

	var c interfaces.Ec2Client
	if i, found := g.ec2Cache.Get(loc.String()); found {
		log.Debug().Msg("Found EC2 client in cache")
		c = i.(*ec2.Client)
		return c, nil
	}

	log.Debug().Msg("EC2 client not found in cache, creating a new one")
	client, err := g.createEc2ClientForLocation(ctx, loc)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create EC2 client")
		return nil, err
	}

	log.Debug().Msg("Returning EC2 client")
	return client, nil
}

func (g *PlatformAwsEcsGateway) createEc2ClientForLocation(ctx context.Context, loc provider.Location) (interfaces.Ec2Client, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting EC2 client for location %s", loc.String())

	log.Debug().Msg("Creating EC2 client")
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(loc.String()), config.WithCredentialsProvider(g.creds))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create EC2 client")
		return nil, err
	}

	log.Debug().Msg("Creating ECS client")
	client := ec2.NewFromConfig(cfg)

	log.Debug().Msg("Caching ECS client")
	g.ec2Cache.Set(loc.String(), &client, cache.DefaultExpiration)

	log.Debug().Msg("Returning ECS client")
	return client, nil
}

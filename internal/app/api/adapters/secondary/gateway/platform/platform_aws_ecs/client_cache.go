package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/patrickmn/go-cache"
)

type ClientCreateFunc[C any, O any] func(aws.Config, ...func(O)) C

func getClientForLocation[C any, O any](ctx context.Context, clientFunc ClientCreateFunc[C, O], ch *cache.Cache, loc provider.Location, creds aws.CredentialsProvider) (C, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting client for location %s", loc.String())

	log.Debug().Msg("Checking if client in cache")
	var c C
	if i, found := ch.Get(loc.String()); found {
		log.Debug().Msg("Found client in cache")
		c = i.(C)
		return c, nil
	}

	log.Debug().Msg("Client not found in cache, creating a new one")
	log.Debug().Msg("Creating AWS config")
	cfg, err := createAwsConfig(ctx, loc, creds)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create AWS config")
		return c, err
	}

	log.Debug().Msg("Creating client")
	client := clientFunc(cfg)
	ch.Set(loc.String(), client, cache.DefaultExpiration)

	log.Debug().Msg("Returning client")
	return client, nil
}

func createAwsConfig(ctx context.Context, loc provider.Location, creds aws.CredentialsProvider) (aws.Config, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating AWS config")

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(loc.String()), config.WithCredentialsProvider(creds))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create cfg")
		return cfg, err
	}

	log.Debug().Msg("Returning AWS config")
	return cfg, nil
}

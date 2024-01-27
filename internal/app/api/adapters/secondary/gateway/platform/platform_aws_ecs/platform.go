package platform_aws_ecs

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/patrickmn/go-cache"
)

const (
	defaultExpiration      = 5 * time.Minute
	defaultCleanUpInterval = 10 * time.Minute
)

type PlatformAwsEcsGateway struct {
	// account  interfaces.AwsAccountClient
	ecsCache *cache.Cache
	creds    *credentials.StaticCredentialsProvider
}

func New(accessKey, secretKey string) gateway.Platform {
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	ecsCache := cache.New(defaultExpiration, defaultCleanUpInterval)

	return &PlatformAwsEcsGateway{
		// account:  acc,
		ecsCache: ecsCache,
		creds:    &creds,
	}
}

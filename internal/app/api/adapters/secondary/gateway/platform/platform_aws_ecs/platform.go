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
	ssmCache *cache.Cache // TODO: consolidate these caches into one
	ecsCache *cache.Cache
	ec2Cache *cache.Cache
	iamCache *cache.Cache
	creds    *credentials.StaticCredentialsProvider
}

func New(accessKey, secretKey string) gateway.Platform {
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	ecsCache := cache.New(defaultExpiration, defaultCleanUpInterval)
	ec2Cache := cache.New(defaultExpiration, defaultCleanUpInterval)
	ssmCache := cache.New(defaultExpiration, defaultCleanUpInterval)
	iamCache := cache.New(defaultExpiration, defaultCleanUpInterval)

	return &PlatformAwsEcsGateway{
		// account:  acc,
		iamCache: iamCache,
		ssmCache: ssmCache,
		ecsCache: ecsCache,
		ec2Cache: ec2Cache,
		creds:    &creds,
	}
}

package platform_aws_ecs

import (
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	platform_aws "github.com/awlsring/texit/internal/pkg/platform/aws"
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
	iamCache *cache.Cache
	*platform_aws.BasePlatformAws
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
		BasePlatformAws: &platform_aws.BasePlatformAws{
			Ec2Cache: ec2Cache,
			Creds:    &creds,
		},
	}
}

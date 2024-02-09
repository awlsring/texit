package platform_aws_ec2

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	platform_aws "github.com/awlsring/texit/internal/pkg/platform/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/patrickmn/go-cache"
)

const (
	// TODO: make these override able in config or via input?
	DefaultInstanceType = "t4g.nano"
	DefaultInstanceArch = "arm64"
)

type PlatformAwsEc2Gateway struct {
	*platform_aws.BasePlatformAws
}

func New(accessKey, secretKey string) gateway.Platform {
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
	ec2Cache := cache.New(platform_aws.DefaultCacheExpiration, platform_aws.DefaultCacheCleanUpInterval)

	return &PlatformAwsEc2Gateway{
		BasePlatformAws: &platform_aws.BasePlatformAws{
			Ec2Cache: ec2Cache,
			Creds:    &creds,
		},
	}
}

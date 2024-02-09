package platform_aws

import (
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/patrickmn/go-cache"
)

// Base struct with common functions for AWS clients
// TODO: this might be kinda pointless
type BasePlatformAws struct {
	Ec2Cache *cache.Cache
	Creds    *credentials.StaticCredentialsProvider
}

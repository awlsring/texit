package setup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func loadAwsConfig(access, secret, region string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}
	if region != "" {
		opts = append(opts, awscfg.WithRegion(region))
	}
	if access != "" && secret != "" {
		creds := credentials.NewStaticCredentialsProvider(access, secret, "")
		opts = append(opts, awscfg.WithCredentialsProvider(creds))
	}
	cfg, err := awscfg.LoadDefaultConfig(context.Background(), opts...)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

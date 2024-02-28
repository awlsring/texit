package config

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

const (
	DefaultAccessKeyEnvVar = "AWS_ACCESS_KEY_ID"
	DefaultSecretKeyEnvVar = "AWS_SECRET_ACCESS_KEY"
	DefaultRegionEnvVar    = "AWS_REGION"
)

var (
	ErrMissingAccessKey = errors.New("missing access key")
	ErrMissingSecretKey = errors.New("missing secret key")
	ErrMissingRegionKey = errors.New("missing region")
)

func getCustomOrDefaultFromEnv(c, d string, e error) (string, error) {
	val := os.Getenv(c)
	if val == "" {
		v := os.Getenv(d)
		if v == "" {
			return "", e
		}
		val = v
	}
	return val, nil
}

func AwsAccessKeyFromEnv(k string) (string, error) {
	return getCustomOrDefaultFromEnv(k, DefaultAccessKeyEnvVar, ErrMissingAccessKey)
}

func SecretKeyFromEnv(k string) (string, error) {
	return getCustomOrDefaultFromEnv(k, DefaultSecretKeyEnvVar, ErrMissingSecretKey)
}

func RegionFromEnv(k string) (string, error) {
	return getCustomOrDefaultFromEnv(k, DefaultRegionEnvVar, ErrMissingRegionKey)
}

func LoadAwsConfig(access, secret, region string) (aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}
	if region != "" {
		opts = append(opts, config.WithRegion(region))
	}
	if access != "" && secret != "" {
		creds := credentials.NewStaticCredentialsProvider(access, secret, "")
		opts = append(opts, config.WithCredentialsProvider(creds))
	}
	cfg, err := config.LoadDefaultConfig(context.Background(), opts...)
	if err != nil {
		return aws.Config{}, err
	}
	return cfg, nil
}

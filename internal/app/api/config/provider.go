package config

import (
	"fmt"
	"os"
)

type ProviderType string

func (t ProviderType) String() string {
	return string(t)
}

const (
	ProviderTypeAwsEcs ProviderType = "aws-ecs"
	ProviderTypeAwsEc2 ProviderType = "aws-ec2"
)

const (
	AwsAccessKeyIdSuffix     = "AWS_ACCESS_KEY_ID"
	AwsSecretAccessKeySuffix = "AWS_SECRET_ACCESS_KEY"
)

var (
	ErrMissingProviderAwsAccessKey = fmt.Errorf("missing provider access key")
	ErrMissingProviderAwsSecretKey = fmt.Errorf("missing provider secret key")
	ErrMissingProviderName         = fmt.Errorf("missing provider name")
)

// Configuration for a provider. Currently only AWS ECS.
type ProviderConfig struct {
	// The type of provider, curretly only aws-ecs
	Type ProviderType `yaml:"type"`
	// The access key for the provider. This is only for AWS types.
	AccessKey string `yaml:"accessKey"`
	// The secret key for the provider. This is only for AWS types.
	SecretKey string `yaml:"secretKey"`
	// The name of the provider.
	Name string `yaml:"name"`
}

func providerSecretEnv(name, suffix string) string {
	return fmt.Sprintf("%s_%s", name, suffix)
}

func (c *ProviderConfig) Validate() error {
	switch c.Type {
	case ProviderTypeAwsEcs, ProviderTypeAwsEc2:
		return c.validateAws()
	default:
		return fmt.Errorf("invalid provider type: %s", c.Type)
	}
}

func (c *ProviderConfig) validateAws() error {
	if c.Name == "" {
		return ErrMissingProviderName
	}

	if c.AccessKey == "" {
		key := os.Getenv(providerSecretEnv(c.Name, AwsAccessKeyIdSuffix))
		if key == "" {
			return ErrMissingProviderAwsAccessKey
		}
		c.AccessKey = key
	}

	if c.SecretKey == "" {
		key := os.Getenv(providerSecretEnv(c.Name, AwsSecretAccessKeySuffix))
		if key == "" {
			return ErrMissingProviderAwsSecretKey
		}
		c.SecretKey = key
	}

	return nil
}

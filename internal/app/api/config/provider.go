package config

import "fmt"

type ProviderType string

func (t ProviderType) String() string {
	return string(t)
}

const (
	ProviderTypeAwsEcs ProviderType = "aws-ecs"
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
	// Whether this provider is the default provider.
	Default bool `yaml:"default"`
}

func (c *ProviderConfig) Validate() error {
	switch c.Type {
	case ProviderTypeAwsEcs:
		return c.validateAwsEcs()
	default:
		return fmt.Errorf("invalid provider type: %s", c.Type)
	}
}

func (c *ProviderConfig) validateAwsEcs() error {
	if c.AccessKey == "" {
		return ErrMissingProviderAwsAccessKey
	}

	if c.SecretKey == "" {
		return ErrMissingProviderAwsSecretKey
	}

	if c.Name == "" {
		return ErrMissingProviderName
	}

	return nil
}

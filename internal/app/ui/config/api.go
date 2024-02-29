package config

import (
	"errors"
	"os"
)

var (
	ErrMissingAddress = errors.New("missing texit address")
	ErrMissingApiKey  = errors.New("missing texit apiKey")
)

const (
	TexitEndpointEnvVar = "TEXIT_ENDPOINT"
	TexitApiKeyEnvVar   = "TEXIT_API_KEY"
)

// Configuration to connect to the API server
type ApiConfig struct {
	Address string `yaml:"address"`
	ApiKey  string `yaml:"apiKey"`
}

func (c *ApiConfig) Validate() error {
	if c.Address == "" {
		k := os.Getenv(TexitEndpointEnvVar)
		if k == "" {
			return ErrMissingAddress
		}
		c.Address = k
	}

	if c.ApiKey == "" {
		k := os.Getenv(TexitApiKeyEnvVar)
		if k == "" {
			return ErrMissingApiKey
		}
		c.ApiKey = k
	}

	return nil
}

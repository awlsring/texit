package config

import "errors"

var (
	ErrMissingAddress = errors.New("missing texit address")
	ErrMissingApiKey  = errors.New("missing texit apiKey")
)

// Configuration to connect to the API server
type ApiConfig struct {
	Address string `yaml:"address"`
	ApiKey  string `yaml:"apiKey"`
}

func (c *ApiConfig) Validate() error {
	if c.Address == "" {
		return ErrMissingAddress
	}

	if c.ApiKey == "" {
		return ErrMissingApiKey
	}

	return nil
}

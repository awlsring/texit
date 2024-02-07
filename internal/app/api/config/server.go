package config

import (
	"math/rand"
	"os"

	"github.com/awlsring/texit/internal/pkg/tsn"
)

const (
	ServerApiKeyEnv = "SERVER_API_KEY"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomApiKey() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Configuration for the server
type ServerConfig struct {
	// The address to listen on
	Address string `yaml:"address"`
	// If specified, the server will join the tailnet with the config provided
	Tailnet *tsn.Config `yaml:"tailnet"`
	// A static key to use for the API. If not specified, one will be generated.
	APIKey string `yaml:"apiKey"` //TODO: make this better one day
}

func (c *ServerConfig) Validate() error {
	if c.Address == "" {
		c.Address = ":7032"
	}

	if c.Tailnet != nil {
		if err := c.Tailnet.Validate(); err != nil {
			return err
		}
	}

	if c.APIKey == "" {
		key := os.Getenv(ServerApiKeyEnv)
		if key == "" {
			c.APIKey = generateRandomApiKey()
		} else {
			c.APIKey = key
		}
	}

	return nil
}

package config

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	DefaultHostName = "texit"
	ServerApiKeyEnv = "SERVER_API_KEY"
)

var (
	ErrMissingTailnetAuthKey = fmt.Errorf("missing tailnet preauth key")
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Configuration for the server tailnet connection
// If this is specified, the server will be started using tsnet and will join the tailnet.
type ServerTailnetConfig struct {
	// Required, the authkey to join the tailnet
	AuthKey string `yaml:"authkey"`
	// The hostname to use on the tailnet
	Hostname string `yaml:"hostname"`
	// The directory to store the state of the tailnet. If not specified, the default will be used.
	StateDir string `yaml:"state"`
	// Whether to use TLS for the tailnet connection
	Tls bool `yaml:"tls"`
	// ControlUrl is the URL of the control server to use. Specify this if you are using Headscale. If not specified, the default tailscale address will be used.
	ControlUrl string `yaml:"controlUrl"`
}

func (c *ServerTailnetConfig) Validate() error {
	if c.AuthKey == "" {
		return ErrMissingTailnetAuthKey
	}

	if c.Hostname == "" {
		c.Hostname = DefaultHostName
	}

	return nil
}

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
	Tailnet *ServerTailnetConfig `yaml:"tailnet"`
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

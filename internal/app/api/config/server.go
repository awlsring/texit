package config

import "fmt"

const (
	DefaultHostName = "texit"
	State
)

var (
	ErrMissingTailnetAuthKey = fmt.Errorf("missing tailnet preauth key")
)

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
	// Tls bool `yaml:"tls"`
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

// Configuration for the server
type ServerConfig struct {
	Address string               `yaml:"address"`
	Tailnet *ServerTailnetConfig `yaml:"tailnet"`
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

	return nil
}

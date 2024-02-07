package config

import (
	"errors"

	"github.com/awlsring/texit/internal/pkg/tsn"
)

const (
	DefaultAddress  = ":8032"
	DefaultHostName = "texit-discord-bot"
)

var (
	ErrMissingTailnetAuthKey = errors.New("missing tailnet preauth key")
)

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
	Funnel     bool   `yaml:"funnel"`
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

type ServerConfig struct {
	Address string      `yaml:"address"`
	Tailnet *tsn.Config `yaml:"tailnet"`
}

func (c *ServerConfig) Validate() error {
	if c.Address == "" {
		c.Address = DefaultAddress
	}

	if c.Tailnet != nil {
		if err := c.Tailnet.Validate(); err != nil {
			return err
		}
	}

	return nil
}

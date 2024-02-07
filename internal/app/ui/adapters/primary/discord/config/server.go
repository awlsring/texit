package config

import (
	"errors"

	"github.com/awlsring/texit/internal/pkg/tsn"
)

const (
	DefaultAddress  = ":8032"
	DefaultHostName = "texit-discord-bot"
	DefaultTsState  = "/var/lib/texit-discord/tsstate"
)

var (
	ErrMissingTailnetAuthKey = errors.New("missing tailnet preauth key")
)

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
		if c.Tailnet.StateDir == "" {
			c.Tailnet.StateDir = DefaultTsState
		}
	}

	return nil
}

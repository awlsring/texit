package config

import (
	"github.com/awlsring/texit/internal/app/ui/config"
)

type Config struct {
	LogLevel     string           `yaml:"logLevel"`
	Api          config.ApiConfig `yaml:"api"`
	Notification NotifierConfig   `yaml:"notifier"`
	Server       ServerConfig     `yaml:"server"`
	Discord      DiscordBotConfig `yaml:"discord"`
	Tracker      *TrackerConfig   `yaml:"tracker"`
}

func (c *Config) Validate() error {
	if c.LogLevel == "" {
		c.LogLevel = "info"
	}

	if c.Tracker == nil {
		c.Tracker = NewDefaultTrackerConfig()
	}
	if err := c.Tracker.Validate(); err != nil {
		return err
	}

	if err := c.Api.Validate(); err != nil {
		return err
	}

	if err := c.Server.Validate(); err != nil {
		return err
	}

	if err := c.Discord.Validate(); err != nil {
		return err
	}

	if err := c.Notification.Validate(); err != nil {
		return err
	}

	return nil
}

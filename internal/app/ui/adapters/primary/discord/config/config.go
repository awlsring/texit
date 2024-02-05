package config

import "github.com/awlsring/texit/internal/app/ui/config"

type Config struct {
	Api     config.ApiConfig `yaml:"api"`
	Server  ServerConfig     `yaml:"server"`
	Discord DiscordBotConfig `yaml:"discord"`
}

func (c Config) Validate() error {
	if err := c.Api.Validate(); err != nil {
		return err
	}

	if err := c.Server.Validate(); err != nil {
		return err
	}

	if err := c.Discord.Validate(); err != nil {
		return err
	}

	return nil
}

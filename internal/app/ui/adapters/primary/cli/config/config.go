package config

import "github.com/awlsring/texit/internal/app/ui/config"

type Config struct {
	Api config.ApiConfig `yaml:"api"`
}

func (c Config) Validate() error {
	if err := c.Api.Validate(); err != nil {
		return err
	}

	return nil
}

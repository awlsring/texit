package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server    ServerConfig     `yaml:"server"`
	Tailnet   TailnetConfig    `yaml:"tailscale"`
	Database  DatabaseConfig   `yaml:"database"`
	Providers []ProviderConfig `yaml:"providers"`
}

// Loads the application config from a file at the specified path.
func LoadFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.Tailnet.Validate()
	if err != nil {
		return nil, err
	}

	err = cfg.Database.Validate()
	if err != nil {
		return nil, err
	}

	err = cfg.Server.Validate()
	if err != nil {
		return nil, err
	}

	for _, p := range cfg.Providers {
		err = p.Validate()
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

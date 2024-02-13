package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server    *ServerConfig     `yaml:"server"`
	Tailnets  []*TailnetConfig  `yaml:"tailnets"`
	Database  *DatabaseConfig   `yaml:"database"`
	Providers []*ProviderConfig `yaml:"providers"`
}

func LoadFromData(data []byte) (*Config, error) {
	var cfg Config
	err := yaml.Unmarshal(data, &cfg)
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

	if len(cfg.Tailnets) == 0 {
		return nil, errors.New("no tailnets configured")
	}
	for _, t := range cfg.Tailnets {
		err = t.Validate()
		if err != nil {
			return nil, err
		}
	}

	if len(cfg.Providers) == 0 {
		return nil, errors.New("no providers configured")
	}
	for _, p := range cfg.Providers {
		err = p.Validate()
		if err != nil {
			return nil, err
		}
	}

	return &cfg, nil
}

// Loads the application config from a file at the specified path.
func LoadFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return LoadFromData(data)
}

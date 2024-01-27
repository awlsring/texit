package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration to connect to the API server
type ApiConfig struct {
	Address string `yaml:"address"`
	ApiKey  string `yaml:"apiKey"`
}

type Config struct {
	Api ApiConfig `yaml:"api"`
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

	// TODO: validate config
	return &cfg, nil
}

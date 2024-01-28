package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration for a provider. Currently only AWS ECS.
type ProviderConfig struct {
	Type      string `yaml:"type"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Name      string `yaml:"name"`
	Default   bool   `yaml:"default"`
}

// Configuration for the server
type ServerConfig struct {
	Address string `yaml:"address"`
}

// Configuration for the database
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Config struct {
	Server    ServerConfig     `yaml:"server"`
	Tailnet   TailnetConfig    `yaml:"tailscale"`
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

	// TODO: validate config
	return &cfg, nil
}

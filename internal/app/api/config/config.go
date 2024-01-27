package config

import (
	"encoding/json"
	"os"
)

type TailnetType string // tailscale or headscale
const (
	TailnetTypeTailscale TailnetType = "tailscale"
	TailnetTypeHeadscale TailnetType = "headscale"
)

// Configuration for the tailnet exit nodes will join
type TailnetConfig struct {
	// The type of tailnet, tailscale or headscale
	Type TailnetType `json:"type"`
	// The network of the tailnet. On tailscale, this is your tailnet name. On headscale, this is the server address.
	Network string `json:"tailnet"`
	// The api token to communicate with the tailnet
	ApiKey string `json:"apiKey"`
}

type ProviderType string

const (
	ProviderTypeAwsEcs ProviderType = "aws:ecs"
)

// Configuration for a provider. Currently only AWS ECS.
type ProviderConfig struct {
	Type      ProviderType `json:"type"`
	AccessKey string       `json:"accessKey"`
	SecretKey string       `json:"secretKey"`
	Region    string       `json:"region"`
	Name      string       `json:"name"`
}

// Configuration for the server
type ServerConfig struct {
	Address string `json:"address"`
}

// Configuration for the database
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	Server    ServerConfig     `json:"server"`
	Tailscale TailnetConfig    `json:"tailscale"`
	Providers []ProviderConfig `json:"providers"`
}

// Loads the application config from a file at the specified path.
func LoadFromFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// TODO: validate config
	return &cfg, nil
}

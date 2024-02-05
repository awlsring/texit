package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config interface {
	Validate() error
}

// Loads the application config from a file at the specified path.
func LoadFromFile[C Config](path string) (*C, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg C
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	// TODO: validate config
	return &cfg, nil
}

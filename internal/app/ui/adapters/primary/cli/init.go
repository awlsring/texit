package cli

import (
	"os"

	cliconfig "github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/config"
	"github.com/awlsring/texit/internal/app/ui/config"
	"gopkg.in/yaml.v2"
)

const (
	configOverrideFlag = "TEXIT_CLI_CONFIG"
	defaultConfigPath  = ".texit/config.yaml"
)

func texitDir() string {
	return os.ExpandEnv("$HOME/" + ".texit")
}

func MakeTexitDir() {
	if _, err := os.Stat(texitDir()); os.IsNotExist(err) {
		err := os.MkdirAll(texitDir(), 0755)
		if err != nil {
			panic(err)
		}
	}
}

func configPath() string {
	return os.ExpandEnv("$HOME/" + defaultConfigPath)
}

func InitDefaultConfig() {
	cfg := cliconfig.Config{
		Api: config.ApiConfig{
			Address: "http://myserver:7032",
			ApiKey:  "changeme",
		},
	}
	path := configPath()
	y, _ := yaml.Marshal(cfg)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.WriteFile(path, y, 0644)
		if err != nil {
			panic(err)
		}
	}
}

func LoadConfig() (*cliconfig.Config, error) {
	configPath := configPath()
	if override := os.Getenv(configOverrideFlag); override != "" {
		configPath = override
	}

	cfg, err := config.LoadFromFile[cliconfig.Config](configPath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

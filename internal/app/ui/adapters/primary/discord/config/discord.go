package config

import "errors"

var (
	ErrMissingApplicationId = errors.New("missing applicationId")
	ErrMissingPublicKey     = errors.New("missing publicKey")
	ErrMissingToken         = errors.New("missing token")
)

type DiscordBotConfig struct {
	ApplicationId string   `yaml:"applicationId"`
	PublicKey     string   `yaml:"publicKey"`
	GuildIds      []string `yaml:"guildIds"`
	Token         string   `yaml:"token"`
	Authorized    []string `yaml:"authorized"`
}

func (c *DiscordBotConfig) Validate() error {
	if c.ApplicationId == "" {
		return ErrMissingApplicationId
	}

	if c.PublicKey == "" {
		return ErrMissingPublicKey
	}

	if c.Token == "" {
		return ErrMissingToken
	}

	return nil
}

package config

import "errors"

var (
	ErrMissingApplicationId = errors.New("missing applicationId")
	ErrMissingPublicKey     = errors.New("missing publicKey")
	ErrMissingToken         = errors.New("missing token")
)

type DiscordBotConfig struct {
	ApplicationId string `yaml:"applicationId"`
	PublicKey     string `yaml:"publicKey"`
	Token         string `yaml:"token"`
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

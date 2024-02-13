package config

import (
	"errors"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/pkg/appinit"
)

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

func (c *DiscordBotConfig) AuthorizedAsSnowflakes() ([]tempest.Snowflake, error) {
	authorized := []tempest.Snowflake{}
	for _, id := range c.Authorized {
		s, err := tempest.StringToSnowflake(id)
		appinit.PanicOnErr(err)
		authorized = append(authorized, s)
	}
	return authorized, nil
}

func (c *DiscordBotConfig) AuthorizedGuildsAsSnowflakes() ([]tempest.Snowflake, error) {
	var guilds []tempest.Snowflake
	guilds = nil
	if len(c.GuildIds) > 0 {
		for _, id := range c.GuildIds {
			s, err := tempest.StringToSnowflake(id)
			appinit.PanicOnErr(err)
			guilds = append(guilds, s)
		}
	}
	return guilds, nil
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

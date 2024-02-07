package discord

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/rs/zerolog"
)

type BotOption func(*Bot)

func WithLogLevel(level zerolog.Level) BotOption {
	return func(b *Bot) {
		b.logLevel = level
	}
}

func WithAuthorizedUsers(users []tempest.Snowflake) BotOption {
	return func(b *Bot) {
		b.authorized = users
	}
}

func WithGuilds(guilds []tempest.Snowflake) BotOption {
	return func(b *Bot) {
		b.guildIds = guilds
	}
}

package discord

import (
	tempest "github.com/Amatsagu/Tempest"
	comctx "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

func (b *Bot) auth(ctx *comctx.CommandContext) bool {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Checking authorization")

	if len(b.authorized) == 0 {
		log.Debug().Msg("No authorized users, allowing")
		return true
	}

	for _, id := range b.authorized {
		if ctx.Requester() == id {
			log.Debug().Msg("User is authorized")
			return true
		}
		if ctx.RequesterRoles() != nil {
			for _, r := range ctx.RequesterRoles() {
				if r == id {
					log.Debug().Msg("User is in authorized role")
					return true
				}
			}
		}
	}

	log.Debug().Msgf("User %d is not authorized", ctx.Requester())
	return false
}

type CommandFunc func(*comctx.CommandContext)

func (b *Bot) CommandPreflight(comFunc CommandFunc) func(itx *tempest.CommandInteraction) {
	return func(itx *tempest.CommandInteraction) {
		ctx, err := comctx.InitContext(b.tmpst, itx, b.logLevel)
		if err != nil {
			return
		}
		if !b.auth(ctx) {
			_ = ctx.SendLinearReply("You are not authorized to use this command", true)
			return
		}
		log := ctx.Logger()
		log.Debug().Msg("Deferring command interaction")
		if err := ctx.DeferResponse(); err != nil {
			log.Error().Err(err).Msg("Failed to defer command interaction")
			if err = ctx.SendLinearReply("Command failed with an unknown error!", true); err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		comFunc(ctx)
	}
}

type AutoCompleteFunc func(*comctx.CommandContext) []tempest.Choice

func (b *Bot) AutoCompletePreflight(logLevel zerolog.Level, comFunc AutoCompleteFunc) func(itx tempest.CommandInteraction) []tempest.Choice {
	return func(itx tempest.CommandInteraction) []tempest.Choice {
		ctx, err := comctx.InitContext(b.tmpst, &itx, logLevel)
		if err != nil {
			return nil
		}
		if !b.auth(ctx) {
			return nil
		}
		return comFunc(ctx)
	}
}

package command

import (
	"context"

	tempest "github.com/Amatsagu/Tempest"
	comctx "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type CommandFunc func(*comctx.CommandContext)

func CommandWrapper(logLevel zerolog.Level, tmpst *tempest.Client, comFunc CommandFunc) func(itx *tempest.CommandInteraction) {
	return func(itx *tempest.CommandInteraction) {
		ctx, err := comctx.InitContext(tmpst, itx, logLevel)
		if err != nil {
			return
		}
		log := ctx.Logger()
		log.Debug().Msg("Deferring command interaction")
		if err := itx.Defer(true); err != nil {
			log.Error().Err(err).Msg("Failed to defer command interaction")
			if err = ctx.SendLinearReply("Command failed with an unknown error!", true); err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		comFunc(ctx)
	}
}

type AutoCompleteFunc func(context.Context, tempest.CommandInteraction) []tempest.Choice

func AutoCompleteWrapper(logLevel zerolog.Level, comFunc AutoCompleteFunc) func(itx tempest.CommandInteraction) []tempest.Choice {
	return func(itx tempest.CommandInteraction) []tempest.Choice {
		ctx := logger.InitContextLogger(context.Background(), logLevel)
		return comFunc(ctx, itx)
	}
}

package handler

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
)

func (h *Handler) ServerHealthCheck(ctx *context.CommandContext) {
	log := ctx.Logger()

	log.Debug().Msg("Checking server health")

	log.Debug().Msg("Calling server health method")
	err := h.apiSvc.CheckServerHealth(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Server is unhealthy")
		if err := ctx.EditResponse(fmt.Sprintf("Server is unhealthy: %s", err.Error())); err != nil {
			log.Error().Err(err).Msg("Failed to write bot response")
		}
		return
	}
	log.Debug().Msg("Server is healthy, writing bot response")
	if err := ctx.EditResponse("Server is healthy!"); err != nil {
		log.Error().Err(err).Msg("Failed to write bot response")
	}
}

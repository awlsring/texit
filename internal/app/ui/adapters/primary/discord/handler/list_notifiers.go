package handler

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
)

func (h *Handler) ListNotifiers(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Listing notifiers")

	log.Debug().Msg("Getting notifiers from service")
	ns, err := h.apiSvc.ListNotifiers(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Error listing notifiers")
		InternalErrorResponse(ctx)
		return
	}

	if len(ns) == 0 {
		log.Debug().Msg("No notifiers found")
		_ = ctx.EditResponse("No notifiers found")
		return
	}

	log.Debug().Msg("notifiers found, creating message")
	msg := "Here's a list of all current known notifiers...\n\n"
	for _, n := range ns {
		msg += fmt.Sprintf("- %s (%s) - Endpoint `%s` \n", n.Name.String(), n.Type.String(), n.Endpoint.String())
	}

	_ = ctx.EditResponse(msg)
}

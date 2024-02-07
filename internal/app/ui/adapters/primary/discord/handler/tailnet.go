package handler

import (
	"errors"
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

func (h *Handler) DescribeTailnet(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Describing tailnet")

	log.Debug().Msg("Getting tailnet name")
	tailnetName, ok := ctx.GetOptionValue(command.OptionTailnetName)
	if !ok {
		log.Error().Msg("Failed to get tailnet name from interaction")
		_ = ctx.EditResponse("Please specify a tailnet name.", true)
		return
	}

	provName, err := tailnet.IdentifierFromString(tailnetName.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet name")
		TailnetNameInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Getting tailnet from service")
	p, err := h.tailSvc.DescribeTailnet(ctx, provName)
	if err != nil {
		if errors.Is(err, service.ErrUnknownTailnet) {
			UnknownTailnetResponse(ctx, provName.String())
			return
		}
		log.Warn().Err(err).Msg("Error getting tailnet")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Creating message")
	msg := fmt.Sprintf("Here's some information about %s...\n\n", p.Name.String())
	msg += fmt.Sprintf("Name: %s\n", p.Name.String())
	msg += fmt.Sprintf("Type: %s\n", p.Type.String())

	_ = ctx.EditResponse(msg, true)
}

func (h *Handler) ListTailnets(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Listing tailnets")

	log.Debug().Msg("Getting tailnets from service")
	ps, err := h.tailSvc.ListTailnets(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Error listing tailnets")
		InternalErrorResponse(ctx)
		return
	}

	if len(ps) == 0 {
		log.Debug().Msg("No tailnets found")
		_ = ctx.EditResponse("No tailnets found", true)
		return
	}

	log.Debug().Msg("Tailnets found, creating message")
	msg := "Here's a list of all current known tailnets...\n\n"
	for _, p := range ps {
		msg += fmt.Sprintf("- %s (%s)\n", p.Name.String(), p.Type.String())
	}

	_ = ctx.EditResponse(msg, true)
}

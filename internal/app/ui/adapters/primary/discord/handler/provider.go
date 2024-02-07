package handler

import (
	"errors"
	"fmt"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

func (h *Handler) DescribeProvider(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Describing provider")

	log.Debug().Msg("Getting provider name")
	providerName, ok := ctx.GetOptionValue(command.OptionProviderName)
	if !ok {
		log.Error().Msg("Failed to get provider name from interaction")
		_ = ctx.EditResponse("Please specify a provider name.", true)
		return
	}

	provName, err := provider.IdentifierFromString(providerName.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		ProviderNameInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Getting provider from service")
	p, err := h.provSvc.DescribeProvider(ctx, provName)
	if err != nil {
		if errors.Is(err, service.ErrUnknownProvider) {
			UnknownProviderResponse(ctx, provName.String())
			return
		}
		log.Warn().Err(err).Msg("Error getting provider")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Creating message")
	msg := fmt.Sprintf("Here's some information about %s...\n\n", p.Name.String())
	msg += fmt.Sprintf("Name: %s\n", p.Name.String())
	msg += fmt.Sprintf("Platform: %s\n", p.Platform.String())

	_ = ctx.EditResponse(msg, true)
}

func (h *Handler) ListProviders(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Listing providers")

	log.Debug().Msg("Getting providers from service")
	ps, err := h.provSvc.ListProviders(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Error listing providers")
		InternalErrorResponse(ctx)
		return
	}

	if len(ps) == 0 {
		log.Debug().Msg("No providers found")
		_ = ctx.EditResponse("No providers found", true)
		return
	}

	log.Debug().Msg("Providers found, creating message")
	msg := "Here's a list of all current known providers...\n\n"
	for _, p := range ps {
		msg += fmt.Sprintf("- %s (%s)\n", p.Name.String(), p.Platform.String())
	}

	_ = ctx.EditResponse(msg, true)
}

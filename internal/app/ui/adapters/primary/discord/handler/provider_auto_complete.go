package handler

import (
	"context"
	"strings"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (h *Handler) ProviderNameAutoComplete(ctx context.Context, itx *tempest.CommandInteraction, name, filter string) []tempest.Choice {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Auto completing provider name")

	log.Debug().Msg("Calling list providers method")
	pro, err := h.provSvc.ListProviders(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list providers")
		return []tempest.Choice{}
	}

	log.Debug().Msg("Building choices")
	choices := []tempest.Choice{}
	for _, p := range pro {
		if !strings.Contains(p.Name.String(), filter) {
			continue
		}
		log.Debug().Str("provider_name", p.Name.String()).Msg("Adding provider to choices")
		choices = append(choices, tempest.Choice{
			Name:  p.Name.String(),
			Value: p.Name.String(),
		})
	}

	log.Debug().Msg("Returning choices")
	return choices
}

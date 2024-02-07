package handler

import (
	"context"
	"fmt"
	"strings"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (h *Handler) NodeIdAutoComplete(ctx context.Context, itx tempest.CommandInteraction, name, filter string) []tempest.Choice {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Auto completing node id")

	log.Debug().Msg("Calling list nodes method")
	pro, err := h.apiSvc.ListNodes(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes")
		return []tempest.Choice{}
	}

	log.Debug().Msg("Building choices")
	choices := []tempest.Choice{}
	for _, p := range pro {
		if !strings.Contains(p.Identifier.String(), filter) {
			continue
		}
		log.Debug().Str("node_id", p.Identifier.String()).Msg("Adding node to choices")
		choices = append(choices, tempest.Choice{
			Name:  fmt.Sprintf("%s (%s)", p.Identifier.String(), p.TailnetName.String()),
			Value: p.Identifier.String(),
		})
	}

	log.Debug().Msg("Returning choices")
	return choices
}

func (h *Handler) ProviderNameAutoComplete(ctx context.Context, itx tempest.CommandInteraction, name, filter string) []tempest.Choice {
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

func (h *Handler) TailnetNameAutoComplete(ctx context.Context, itx tempest.CommandInteraction, name, filter string) []tempest.Choice {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Auto completing tailnet name")

	log.Debug().Msg("Calling list tailnets method")
	pro, err := h.tailSvc.ListTailnets(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list tailnets")
		return []tempest.Choice{}
	}

	log.Debug().Msg("Building choices")
	choices := []tempest.Choice{}
	for _, p := range pro {
		if !strings.Contains(p.Name.String(), filter) {
			continue
		}
		log.Debug().Str("tailnet_name", p.Name.String()).Msg("Adding tailnet to choices")
		choices = append(choices, tempest.Choice{
			Name:  p.Name.String(),
			Value: p.Name.String(),
		})
	}

	log.Debug().Msg("Returning choices")
	return choices
}

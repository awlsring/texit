package handler

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/embed"
)

func (h *Handler) ListNodes(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Listing nodes")

	log.Debug().Msg("Getting nodes from service")
	nodes, err := h.nodeSvc.ListNodes(ctx)
	if err != nil {
		log.Warn().Err(err).Msg("Error listing nodes")
		InternalErrorResponse(ctx)
		return
	}

	if len(nodes) == 0 {
		log.Debug().Msg("No nodes found")
		_ = ctx.EditResponse("No nodes found", true)
		return
	}

	log.Debug().Msg("Nodes found, creating embeds")
	ems := []*tempest.Embed{}
	for _, n := range nodes {
		em := embed.NodeAsEmbed(n)
		ems = append(ems, em)
	}

	log.Debug().Msg("Sending embeds in message")
	_ = ctx.EditReply(tempest.ResponseMessageData{
		Content: "Here's a list of all current known nodes...",
		Embeds:  ems,
	}, true)
}

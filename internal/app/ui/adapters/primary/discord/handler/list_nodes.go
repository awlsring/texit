package handler

import (
	"fmt"
	"time"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
)

func (h *Handler) ListNodes(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("Listing nodes")

	nodes, err := h.apiSvc.ListNodes(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error listing nodes")
		_ = ctx.EditResponse(fmt.Sprintf("Error listing nodes. Error: %s", err.Error()), true)
		return
	}

	if len(nodes) == 0 {
		_ = ctx.EditResponse("No nodes found", true)
		return
	}

	msg := "Nodes:\n"
	msg += "-------------------------------------------\n"
	for _, n := range nodes {
		msg += fmt.Sprintf("Node ID: `%s`\n", n.Identifier.String())
		msg += fmt.Sprintf("Tailnet Name: `%s`\n", n.TailnetName.String())
		msg += fmt.Sprintf("Provider: `%s`\n", n.Provider.String())
		msg += fmt.Sprintf("Tailnet: `%s`\n", n.Tailnet.String())
		msg += fmt.Sprintf("Location: `%s`\n", n.Location.String())
		msg += fmt.Sprintf("Ephemeral: `%t`\n", n.Ephemeral)
		msg += fmt.Sprintf("Created At: `%s`\n", n.CreatedAt.Format(time.RFC1123))
		msg += fmt.Sprintf("Updated At: `%s`\n", n.UpdatedAt.Format(time.RFC1123))
		msg += "-------------------------------------------\n"

	}
	_ = ctx.EditResponse(msg, true)
}

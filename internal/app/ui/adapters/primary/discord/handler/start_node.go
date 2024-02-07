package handler

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (h *Handler) StartNode(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("starting node handler")

	nodeIdStr, ok := ctx.GetOptionValue(option.NodeId)
	if !ok {
		log.Error().Msg("Failed to get node ID from interaction")
		_ = ctx.EditResponse("Please specify an node id to start", true)
		return
	}
	nodeId, err := node.IdentifierFromString(nodeIdStr.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node ID")
		_ = ctx.EditResponse(fmt.Sprintf("Failed to parse to provided node id. Error: %s", err.Error()), true)
		return
	}

	log.Debug().Msg("starting node from service")
	err = h.apiSvc.StartNode(ctx, nodeId)
	if err != nil {
		log.Warn().Err(err).Msg("Error starting node")
		_ = ctx.EditResponse("Error starting node", true)
		return
	}

	log.Debug().Msg("Sending bot response")
	_ = ctx.EditResponse("The node is starting!", true)
}

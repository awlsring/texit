package handler

import (
	"errors"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (h *Handler) StopNode(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("stopping node handler")

	nodeIdStr, ok := ctx.GetOptionValue(command.OptionNodeId)
	if !ok {
		log.Error().Msg("Failed to get node ID from interaction")
		_ = ctx.EditResponse("Please specify an node id to stop", true)
		return
	}
	nodeId, err := node.IdentifierFromString(nodeIdStr.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node ID")
		NodeIdInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("stopping node from service")
	err = h.apiSvc.StopNode(ctx, nodeId)
	if err != nil {
		if errors.Is(err, service.ErrUnknownNode) {
			UnknownNodeResponse(ctx, nodeId.String())
			return
		}
		log.Warn().Err(err).Msg("Error stopping node")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Sending bot response")
	_ = ctx.EditResponse("The node is now stopping!", true)
}

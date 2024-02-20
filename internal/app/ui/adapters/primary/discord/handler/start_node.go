package handler

import (
	"errors"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (h *Handler) StartNode(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("starting node handler")

	nodeIdStr, ok := ctx.GetOptionValue(command.OptionNodeId)
	if !ok {
		log.Error().Msg("Failed to get node ID from interaction")
		_ = ctx.EditResponse("Please specify an node id to start")
		return
	}
	nodeId, err := node.IdentifierFromString(nodeIdStr.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node ID")
		NodeIdInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("starting node from service")
	err = h.apiSvc.StartNode(ctx, nodeId)
	if err != nil {
		if errors.Is(err, service.ErrUnknownNode) {
			UnknownNodeResponse(ctx, nodeId.String())
			return
		}
		if errors.Is(err, service.ErrInvalidInputError) {
			InvalidInputErrorResponse(ctx, err)
			return
		}
		log.Warn().Err(err).Msg("Error starting node")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Sending bot response")
	_ = ctx.EditResponse("The node is starting!")
}

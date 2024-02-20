package handler

import (
	"errors"
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (h *Handler) DeprovisionNode(ctx *context.CommandContext) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deprovisioning node")

	log.Debug().Msg("Getting node id")
	nodeId, ok := ctx.GetOptionValue(command.OptionNodeId)
	if !ok {
		log.Error().Msg("Failed to get node id from interaction")
		_ = ctx.EditResponse("Please specify a node id.")
		return
	}
	n, err := node.IdentifierFromString(nodeId.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		NodeIdInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Calling deprovision node method")
	exId, err := h.apiSvc.DeprovisionNode(ctx, n)
	if err != nil {
		if errors.Is(err, service.ErrUnknownNode) {
			UnknownNodeResponse(ctx, n.String())
			return
		}
		if errors.Is(err, service.ErrInvalidInputError) {
			InvalidInputErrorResponse(ctx, err)
			return
		}
		log.Warn().Err(err).Msg("Error deprovisioning node")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Deprovision node workflow started, writing bot response")
	if err = ctx.EditResponse(fmt.Sprintf("Deprovision node workflow started. The execution id is %s\n\nYou'll be sent a message when its finished! This usually takes a few seconds.", fmt.Sprintf("`%s`", exId.String()))); err != nil {
		log.Error().Err(err).Msg("Failed to write bot response")
	}

	log.Debug().Msg("Tracking execution")
	err = h.tracker.AddExecution(ctx, exId.String(), ctx.Requester())
	if err != nil {
		log.Error().Err(err).Msg("Failed to track execution")
	}
}

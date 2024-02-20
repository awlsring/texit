package handler

import (
	"errors"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/embed"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (h *Handler) DescribeNode(ctx *context.CommandContext) {
	log := ctx.Logger()
	log.Debug().Msg("describing node")

	nodeIdStr, ok := ctx.GetOptionValue(command.OptionNodeId)
	if !ok {
		log.Error().Msg("Failed to get node ID from interaction")
		_ = ctx.EditResponse("Please specify an node id to describe")
		return
	}
	nodeId, err := node.IdentifierFromString(nodeIdStr.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node ID")
		NodeIdInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Getting node from service")
	n, err := h.nodeSvc.DescribeNode(ctx, nodeId)
	if err != nil {
		if errors.Is(err, service.ErrUnknownNode) {
			UnknownNodeResponse(ctx, nodeId.String())
			return
		}
		log.Warn().Err(err).Msg("Error getting node")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Node found, creating embed")
	em := embed.NodeAsEmbed(n)

	log.Debug().Msg("Sending embed in message")
	_ = ctx.EditReply(tempest.ResponseMessageData{
		Content: "Here a summary of that node.",
		Embeds:  []*tempest.Embed{em},
	}, true)
}

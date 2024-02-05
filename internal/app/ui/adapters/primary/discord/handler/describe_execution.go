package handler

import (
	"fmt"
	"strings"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (h *Handler) DescribeExecution(ctx *context.CommandContext) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting execution")

	exIdStr, ok := ctx.GetOptionValue(option.ExecutionId)
	if !ok {
		log.Error().Msg("Failed to get execution ID from interaction")
		ctx.EditResponse("Please specify an execution id to describe", true)
		return
	}
	exId, err := workflow.ExecutionIdentifierFromString(exIdStr.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse execution ID")
		ctx.EditResponse(fmt.Sprintf("Failed to parse to provided execution id. Error: %s", err.Error()), true)
		return
	}

	log.Debug().Msg("Calling server health method")
	ex, err := h.apiSvc.GetExecution(ctx, exId)
	if err != nil {
		log.Error().Err(err).Msg("Error describing execution")
		ctx.EditResponse(fmt.Sprintf("Error describing execution: %s", err.Error()), true)
		return
	}
	log.Debug().Msg("Got execution, writing bot response")

	msg := fmt.Sprintf(`### Execution %s
**Workflow**: %s
**Status**: %s
`, ex.Identifier.String(), ex.Workflow.String(), ex.Status.String())
	if ex.Status == workflow.StatusFailed {
		msg += fmt.Sprintf("**Error**: %s", strings.Join(ex.Results, "\n"))
	}
	if ex.Status == workflow.StatusComplete {
		msg += fmt.Sprintf("**Result**: %s", strings.Join(ex.Results, "\n"))
	}

	ctx.EditResponse(msg, true)
}

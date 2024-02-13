package handler

import (
	"fmt"
	"strings"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (h *Handler) DescribeExecution(ctx *context.CommandContext) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting execution")

	exIdStr, ok := ctx.GetOptionValue(command.OptionExecutionId)
	if !ok {
		log.Error().Msg("Failed to get execution ID from interaction")
		_ = ctx.EditResponse("Please specify an execution id to describe", true)
		return
	}
	exId, err := workflow.ExecutionIdentifierFromString(exIdStr.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse execution ID")
		ExecutionIdInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Calling get execution method")
	ex, err := h.apiSvc.GetExecution(ctx, exId)
	if err != nil {
		log.Error().Err(err).Msg("Error getting execution")
		WriteErrorResponse(ctx, err, exId.String())
		return
	}
	log.Debug().Msg("Got execution, writing bot response")
	var msg string
	switch ex.Workflow {
	case workflow.WorkflowNameProvisionNode:
		msg, err = writeProvisionNodeExecutionSummary(ctx, ex)
	case workflow.WorkflowNameDeprovisionNode:
		msg, err = writeDeprovisionNodeExecutionSummary(ctx, ex)
	}
	if err != nil {
		log.Error().Err(err).Msg("Error writing bot response")
		InternalErrorResponse(ctx)
		return
	}

	_ = ctx.EditResponse(msg, true)
}

func writeProvisionNodeExecutionSummary(ctx *context.CommandContext, ex *workflow.Execution) (string, error) {
	output, err := workflow.DeserializeExecutionResult[workflow.ProvisionNodeExecutionResult](ex.Results)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("### Execution %s\n**Workflow**: %s\n**Status**: %s\n", ex.Identifier.String(), ex.Workflow.String(), ex.Status.String())
	if ex.Status == workflow.StatusFailed {
		msg += fmt.Sprintf("**Error**: %s", output.GetError())
	}
	if ex.Status == workflow.StatusComplete {
		msg += fmt.Sprintf("**Node**: %s", output.GetNode())
	}
	return msg, nil
}

func writeDeprovisionNodeExecutionSummary(ctx *context.CommandContext, ex *workflow.Execution) (string, error) {
	output, err := workflow.DeserializeExecutionResult[workflow.DeprovisionNodeExecutionResult](ex.Results)
	if err != nil {
		return "", err
	}

	msg := fmt.Sprintf("### Execution %s\n**Workflow**: %s\n**Status**: %s\n", ex.Identifier.String(), ex.Workflow.String(), ex.Status.String())
	if ex.Status == workflow.StatusFailed {
		msg += fmt.Sprintf("**Error**: %s", output.GetError())
	}
	if len(output.ResourcesFailedToDelete) > 0 {
		msg += fmt.Sprintf("**Resources Failed to Delete**: %s", strings.Join(output.ResourcesFailedToDelete, ","))
	}
	return msg, nil
}

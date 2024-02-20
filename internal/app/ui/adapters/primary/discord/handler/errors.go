package handler

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
)

func WriteErrorResponse(ctx *context.CommandContext, err error, msg string) {
	switch err {
	case service.ErrUnknownNode:
		UnknownNodeResponse(ctx, msg)
	case service.ErrUnknownProvider:
		UnknownProviderResponse(ctx, msg)
	case service.ErrUnknownTailnet:
		UnknownTailnetResponse(ctx, msg)
	case service.ErrUnknownExecution:
		UnknownExecutionResponse(ctx, msg)
	case service.ErrInvalidInputError:
		InvalidInputErrorResponse(ctx, err)
	default:
		InternalErrorResponse(ctx)
	}
}

func UnknownNodeResponse(ctx *context.CommandContext, n string) {
	_ = ctx.EditResponse(fmt.Sprintf("The node `%s` isn't found", n))
}

func UnknownProviderResponse(ctx *context.CommandContext, p string) {
	_ = ctx.EditResponse(fmt.Sprintf("The provider `%s` isn't found", p))
}

func UnknownTailnetResponse(ctx *context.CommandContext, t string) {
	_ = ctx.EditResponse(fmt.Sprintf("The tailnet `%s` isn't found", t))
}

func UnknownExecutionResponse(ctx *context.CommandContext, e string) {
	_ = ctx.EditResponse(fmt.Sprintf("The execution `%s` isn't found", e))
}

func InvalidInputErrorResponse(ctx *context.CommandContext, err error) {
	_ = ctx.EditResponse("Your request has invalid input, double check your inputs.")
}

func NodeIdInvalidConstraintsResponse(ctx *context.CommandContext) {
	_ = ctx.EditResponse("The node id specified doesn't match the required constraints.")
}

func ProviderNameInvalidConstraintsResponse(ctx *context.CommandContext) {
	_ = ctx.EditResponse("The provider name specified doesn't match the required constraints.")
}

func TailnetNameInvalidConstraintsResponse(ctx *context.CommandContext) {
	_ = ctx.EditResponse("The tailnet name specified doesn't match the required constraints.")
}

func ExecutionIdInvalidConstraintsResponse(ctx *context.CommandContext) {
	_ = ctx.EditResponse("The execution id specified doesn't match the required constraints.")
}

func InternalErrorResponse(ctx *context.CommandContext) {
	_ = ctx.EditResponse("An internal error occurred, check the logs for more details.")
}

func ExecutionInternalErrorResponse(ctx *context.CommandContext) {
	_, _ = ctx.SendRequesterPrivateMessage("An internal error occurred running the workflow, check the logs for more details.")
}

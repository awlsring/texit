package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

const (
	pollAmount = 120
	pollDelay  = 5
)

func (h *Handler) ProvisionNode(ctx *context.CommandContext) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Provisioning node")

	log.Debug().Msg("Getting provider name")
	providerName, ok := ctx.GetOptionValue(command.OptionProviderName)
	if !ok {
		log.Error().Msg("Failed to get provider name from interaction")
		_ = ctx.EditResponse("Please specify a provider name.", true)
		return
	}
	pr, err := provider.IdentifierFromString(providerName.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		ProviderNameInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Getting tailnet name")
	tailnetName, ok := ctx.GetOptionValue(command.OptionTailnetName)
	if !ok {
		log.Error().Msg("Failed to get tailnet name from interaction")
		_ = ctx.EditResponse("Please specify a tailnet name.", true)
		return
	}
	tn, err := tailnet.IdentifierFromString(tailnetName.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet name")
		TailnetNameInvalidConstraintsResponse(ctx)
		return
	}

	log.Debug().Msg("Getting provider location")
	providerLocation, ok := ctx.GetOptionValue(command.OptionProviderLocation)
	if !ok {
		log.Error().Msg("Failed to get provider location from interaction")
		_ = ctx.EditResponse("Please specify a provider location.", true)
		return
	}
	pl := provider.Location(providerLocation.(string))

	ephemeral := false
	ephRaw, ok := ctx.GetOptionValue(command.OptionEphemeral)
	if ok {
		ephemeral = ephRaw.(bool)
	}

	log.Debug().Msg("Calling provision node method")
	exId, err := h.apiSvc.ProvisionNode(ctx, pr, pl, tn, ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Error provisioning node")
		InternalErrorResponse(ctx)
		return
	}

	log.Debug().Msg("Provisioned node, writing bot response")
	if err = ctx.EditResponse(fmt.Sprintf("Provision node workflow started. The execution id is %s.\n\nYou'll be sent a message when its ready! This usually takes a few minutes.", fmt.Sprintf("`%s`", exId.String())), true); err != nil {
		log.Error().Err(err).Msg("Failed to write bot response")
	}

	log.Debug().Msg("Polling execution")
	for i := 0; i < pollAmount; i++ {
		log.Debug().Int("poll_count", i).Msg("Polling execution")
		ex, err := h.apiSvc.GetExecution(ctx, exId)
		if err != nil {
			log.Error().Err(err).Msg("Error polling execution")
			ExecutionInternalErrorResponse(ctx)
			return
		}
		log.Debug().Interface("execution", ex).Msg("Execution")
		output, err := workflow.DeserializeExecutionResult[workflow.ProvisionNodeExecutionResult](ex.Results)
		if err != nil {
			log.Error().Err(err).Msg("Error polling execution")
			ExecutionInternalErrorResponse(ctx)
			return
		}
		if ex.Status == workflow.StatusComplete {
			predicatedTailId := fmt.Sprintf("%s-%s", pl, output.GetNode())
			log.Debug().Msg("Execution is complete, writing bot response")
			_, err = ctx.SendRequesterPrivateMessage(fmt.Sprintf("The provision node workflow you requested has completed succesfully.\n\nThe id of your new node is `%s`. (It should appear as something like `%s` on your tailnet)", output.GetNode(), predicatedTailId))
			if err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		if ex.Status == workflow.StatusFailed {
			log.Debug().Msg("Execution is failed, writing bot response")
			_, err = ctx.SendRequesterPrivateMessage(fmt.Sprintf("The provision node workflow you request failed :(.\n\nIt failed on step %s\nErrors encountered: %s", output.GetFailedStep(), strings.Join(output.Errors, ", ")))
			if err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		log.Debug().Msg("Execution is not complete, waiting")
		time.Sleep(pollDelay * time.Second)
	}

}

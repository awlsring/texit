package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
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
	providerName, ok := ctx.GetOptionValue(option.ProviderName)
	if !ok {
		log.Error().Msg("Failed to get provider name from interaction")
		ctx.EditResponse("Please specify a provider name.", true)
		return
	}
	pr, err := provider.IdentifierFromString(providerName.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		ctx.EditResponse("Failed to parse provider name", true)
		return
	}

	log.Debug().Msg("Getting tailnet name")
	tailnetName, ok := ctx.GetOptionValue(option.TailnetName)
	if !ok {
		log.Error().Msg("Failed to get tailnet name from interaction")
		ctx.EditResponse("Please specify a tailnet name.", true)
		return
	}
	tn, err := tailnet.IdentifierFromString(tailnetName.(string))
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet name")
		ctx.EditResponse("Failed to parse tailnet name", true)
		return
	}

	log.Debug().Msg("Getting provider location")
	providerLocation, ok := ctx.GetOptionValue(option.ProviderLocation)
	if !ok {
		log.Error().Msg("Failed to get provider location from interaction")
		ctx.EditResponse("Please specify a provider location.", true)
		return
	}
	pl := provider.Location(providerLocation.(string))

	ephemeral := false
	ephRaw, ok := ctx.GetOptionValue(option.Ephemeral)
	if ok {
		ephemeral = ephRaw.(bool)
	}

	log.Debug().Msg("Calling provision node method")
	exId, err := h.apiSvc.ProvisionNode(ctx, pr, pl, tn, ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Error provisioning node")
		ctx.EditResponse("Error provisioning node", true)
		return
	}

	log.Debug().Msg("Provisioned node, writing bot response")
	if err = ctx.EditResponse(fmt.Sprintf(`Provision node workflow started. The execution id is %s

You'll be sent a message when its ready! This usually takes a few minutes.`, fmt.Sprintf("`%s`", exId.String())), true); err != nil {
		log.Error().Err(err).Msg("Failed to write bot response")
	}

	log.Debug().Msg("Polling execution")
	for i := 0; i < pollAmount; i++ {
		log.Debug().Int("poll_count", i).Msg("Polling execution")
		ex, err := h.apiSvc.GetExecution(ctx, exId)
		if err != nil {
			log.Error().Err(err).Msg("Error polling execution")
			ctx.EditResponse("Error polling execution", true)
			return
		}
		if ex.Status == workflow.StatusComplete {
			log.Debug().Msg("Execution is complete, writing bot response")
			_, err = ctx.SendRequesterPrivateMessage(fmt.Sprintf(`The provision node workflow you requested has completed succesfully.
			
Results: %s`, strings.Join(ex.Results, ", ")))
			if err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		if ex.Status == workflow.StatusFailed {
			log.Debug().Msg("Execution is failed, writing bot response")
			_, err = ctx.SendRequesterPrivateMessage(fmt.Sprintf(`The provision node workflow you request failed :(.
			
Heres the errors: %s`, strings.Join(ex.Results, ", ")))
			if err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		log.Debug().Msg("Execution is not complete, waiting")
		time.Sleep(pollDelay * time.Second)
	}

}

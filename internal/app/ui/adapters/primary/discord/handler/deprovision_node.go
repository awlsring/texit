package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

const (
	deprovisionPollAmount = 20
	deprovisionPollDelay  = 1
)

func (h *Handler) DeprovisionNode(ctx *context.CommandContext) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deprovisioning node")

	log.Debug().Msg("Getting node id")
	nodeId, ok := ctx.GetOptionValue(command.OptionNodeId)
	if !ok {
		log.Error().Msg("Failed to get node id from interaction")
		_ = ctx.EditResponse("Please specify a node id.", true)
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
		WriteErrorResponse(ctx, err, n.String())
		return
	}

	log.Debug().Msg("Deprovision node workflow started, writing bot response")
	if err = ctx.EditResponse(fmt.Sprintf("Deprovision node workflow started. The execution id is %s\n\nYou'll be sent a message when its finished! This usually takes a few seconds.", fmt.Sprintf("`%s`", exId.String())), true); err != nil {
		log.Error().Err(err).Msg("Failed to write bot response")
	}

	log.Debug().Msg("Polling execution")
	for i := 0; i < deprovisionPollAmount; i++ {
		log.Debug().Int("poll_count", i).Msg("Polling execution")
		ex, err := h.apiSvc.GetExecution(ctx, exId)
		if err != nil {
			log.Error().Err(err).Msg("Error polling execution")
			ExecutionInternalErrorResponse(ctx)
			return
		}
		output, err := workflow.DeserializeExecutionResult[workflow.DeprovisionNodeExecutionResult](ex.Results)
		if err != nil {
			log.Error().Err(err).Msg("Error polling execution")
			ExecutionInternalErrorResponse(ctx)
			return
		}
		if ex.Status == workflow.StatusComplete {
			log.Debug().Msg("Execution is complete, writing bot response")
			_, err = ctx.SendRequesterPrivateMessage("The deprovision node workflow you requested has completed succesfully")
			if err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		if ex.Status == workflow.StatusFailed {

			log.Debug().Msg("Execution is failed, writing bot response")
			_, err = ctx.SendRequesterPrivateMessage(fmt.Sprintf("The deprovision node workflow you request failed :(\n\nIt failed on step %s\nErrors: %s", output.GetFailedStep(), strings.Join(output.Errors, ", ")))
			if err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		log.Debug().Msg("Execution is not complete, waiting")
		time.Sleep(deprovisionPollDelay * time.Second)
	}

}

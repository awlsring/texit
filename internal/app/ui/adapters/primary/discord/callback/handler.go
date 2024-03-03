package callback

import (
	"context"
	"fmt"

	tempest "github.com/Amatsagu/Tempest"
	pending_execution "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/execution"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type CallbackHandler struct {
	t  *tempest.Client
	pe pending_execution.Tracker
}

func (h *CallbackHandler) SendDeprovisionFollowUp(ctx context.Context, msg notification.ExecutionMessage) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Sending deprovision follow up")

	output, err := workflow.DeserializeExecutionResult[workflow.DeprovisionNodeExecutionResult](workflow.SerializedExecutionResult(msg.Results))
	if err != nil {
		log.Error().Err(err).Msg("Error getting execution results")
		return
	}

	log.Debug().Msg("Getting user")
	user, err := h.pe.GetExecution(ctx, msg.ExecutionId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return
	}
	defer func() {
		err := h.pe.RemoveExecution(ctx, msg.ExecutionId)
		log.Error().Err(err).Msg("Failed to remove execution")
	}()
	log.Debug().Msg("Forming bot response")

	m := tempest.Message{}

	if msg.Status == workflow.StatusComplete.String() {
		m.Content = "The deprovision node workflow you requested has completed successfully"
	}
	if msg.Status == workflow.StatusFailed.String() {
		m.Content = fmt.Sprintf("The deprovision node workflow you request failed.\n\nError: %s", output.GetError())
	}

	log.Debug().Msg("Sending message")
	_, err = h.t.SendPrivateMessage(user, m)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
	}
}

func (h *CallbackHandler) SendProvisionFollowUp(ctx context.Context, msg notification.ExecutionMessage) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Sending provision follow up")

	output, err := workflow.DeserializeExecutionResult[workflow.ProvisionNodeExecutionResult](workflow.SerializedExecutionResult(msg.Results))
	if err != nil {
		log.Error().Err(err).Msg("Error getting execution results")
		return
	}

	log.Debug().Msg("Getting user")
	user, err := h.pe.GetExecution(ctx, msg.ExecutionId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return
	}
	defer func() {
		err := h.pe.RemoveExecution(ctx, msg.ExecutionId)
		log.Error().Err(err).Msg("Failed to remove execution")
	}()
	log.Debug().Msg("Forming bot response")

	m := tempest.Message{}

	if msg.Status == workflow.StatusComplete.String() {
		m.Content = fmt.Sprintf("The provision node workflow you requested has completed succesfully.\n\nThe id of your new node is `%s`. (It should appear as something like `<location>-%s` on your tailnet)", output.GetNode(), output.GetNode())
	}
	if msg.Status == workflow.StatusFailed.String() {
		m.Content = fmt.Sprintf("The provision node workflow you request failed.\nError: %s", output.GetError())
	}

	log.Debug().Msg("Sending message")
	_, err = h.t.SendPrivateMessage(user, m)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message")
	}
}

func (h *CallbackHandler) Handle(ctx context.Context, msg notification.ExecutionMessage) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Received message")

	switch msg.WorkflowName {
	case workflow.WorkflowNameDeprovisionNode.String():
		h.SendDeprovisionFollowUp(ctx, msg)
	case workflow.WorkflowNameProvisionNode.String():
		h.SendProvisionFollowUp(ctx, msg)
	default:
		log.Error().Str("workflow_name", msg.WorkflowName).Msg("Unknown workflow")
	}

}

func NewCallbackHandler(t *tempest.Client, pe pending_execution.Tracker) *CallbackHandler {
	return &CallbackHandler{
		t:  t,
		pe: pe,
	}
}

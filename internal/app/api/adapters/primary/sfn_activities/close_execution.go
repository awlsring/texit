package sfn_activities

import (
	"context"
	"encoding/json"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type CloseExecutionInput struct {
	WorkflowName string      `json:"workflowName"`
	ExecutionId  string      `json:"executionId"`
	Status       string      `json:"status"`
	Results      interface{} `json:"results"`
	Error        string      `json:"error"`
}

func (h *SfnActivityHandler) closeExecutionActivity(ctx context.Context, input *CloseExecutionInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Closing execution request")

	wf, err := workflow.WorkflowNameFromString(input.WorkflowName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse workflow name")
		return err
	}

	resultsRaw, err := json.Marshal(input.Results)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal results")
		return err
	}

	log.Debug().Msg("Parsing execution id")
	executionId, err := workflow.ExecutionIdentifierFromString(input.ExecutionId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse execution id")
		return err
	}

	status, err := workflow.StatusFromString(input.Status)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse status")
		return err
	}

	var res workflow.ExecutionResult
	switch wf {
	case workflow.WorkflowNameProvisionNode:
		r, err := workflow.DeserializeExecutionResult[workflow.ProvisionNodeExecutionResult](workflow.SerializedExecutionResult(resultsRaw))
		if err != nil {
			log.Error().Err(err).Msg("Failed to deserialize results")
			return err
		}
		if input.Error != "" {
			r.SetError(input.Error)
		}
		res = r
	case workflow.WorkflowNameDeprovisionNode:
		r, err := workflow.DeserializeExecutionResult[workflow.DeprovisionNodeExecutionResult](workflow.SerializedExecutionResult(resultsRaw))
		if err != nil {
			log.Error().Err(err).Msg("Failed to deserialize results")
			return err
		}
		if input.Error != "" {
			r.SetError(input.Error)
		}
	}

	log.Debug().Msg("Closing execution")
	err = h.actSvc.CloseExecution(ctx, executionId, status, res)
	if err != nil {
		log.Error().Err(err).Msg("Failed to close execution")
		return err
	}

	log.Debug().Msg("Signaling execution compplete")
	err = h.notSvc.NotifyExecutionCompletion(ctx, executionId, wf, status, res)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to signal execution complete")
	}

	return nil
}

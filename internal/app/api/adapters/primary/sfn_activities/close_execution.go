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

func serializeResults(results interface{}) (workflow.SerializedExecutionResult, error) {
	if results == nil {
		return workflow.SerializedExecutionResult(""), nil
	}
	resultsRaw, err := json.Marshal(results)
	if err != nil {
		return "", err
	}
	return workflow.SerializedExecutionResult(resultsRaw), nil
}

func (h *SfnActivityHandler) closeExecutionActivity(ctx context.Context, input *CloseExecutionInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Interface("input", input).Msg("Closing execution request")

	wf, err := workflow.WorkflowNameFromString(input.WorkflowName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse workflow name")
		return err
	}

	log.Debug().Msg("Marshalling results")
	resultsRaw, err := serializeResults(input.Results)
	if err != nil {
		log.Error().Err(err).Msg("Failed to marshal results")
		return err
	}
	log.Debug().Interface("results", resultsRaw).Msg("Results marshalled")

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
		r, err := workflow.DeserializeExecutionResult[workflow.ProvisionNodeExecutionResult](resultsRaw)
		if err != nil {
			log.Error().Err(err).Msg("Failed to deserialize results")
			return err
		}
		if input.Error != "" {
			r.SetError(input.Error)
		}
		res = r
	case workflow.WorkflowNameDeprovisionNode:
		r, err := workflow.DeserializeExecutionResult[workflow.DeprovisionNodeExecutionResult](resultsRaw)
		if err != nil {
			log.Error().Err(err).Msg("Failed to deserialize results")
			return err
		}
		if input.Error != "" {
			r.SetError(input.Error)
		}
		res = r
	}
	log.Debug().Interface("results", res).Msg("Results deserialized")

	log.Debug().Msg("Closing execution")
	err = h.actSvc.CloseExecution(ctx, executionId, status, res)
	if err != nil {
		log.Error().Err(err).Msg("Failed to close execution")
		return err
	}

	log.Debug().Msg("Signaling execution complete")
	err = h.notSvc.NotifyExecutionCompletion(ctx, executionId, wf, status, res)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to signal execution complete")
	}

	return nil
}

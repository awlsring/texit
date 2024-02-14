package notification

import (
	"encoding/json"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

type ExecutionMessage struct {
	WorkflowName string `json:"workflowName"`
	ExecutionId  string `json:"executionId"`
	Status       string `json:"status"`
	Results      string `json:"results"`
}

func (e ExecutionMessage) Serialize() (string, error) {
	j, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func NewExecutionMessage(e workflow.ExecutionIdentifier, w workflow.WorkflowName, status workflow.Status, msg, results string) ExecutionMessage {
	return ExecutionMessage{
		WorkflowName: w.String(),
		ExecutionId:  e.String(),
		Status:       status.String(),
		Results:      results,
	}
}

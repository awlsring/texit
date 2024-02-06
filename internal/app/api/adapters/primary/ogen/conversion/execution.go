package conversion

import (
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func TranslateExecutionStatus(t workflow.Status) texit.ExecutionStatus {
	switch t {
	case workflow.StatusPending:
		return texit.ExecutionStatusPending
	case workflow.StatusRunning:
		return texit.ExecutionStatusRunning
	case workflow.StatusComplete:
		return texit.ExecutionStatusCompleted
	case workflow.StatusFailed:
		return texit.ExecutionStatusFailed
	default:
		return texit.ExecutionStatusUnknown
	}
}

func TranslateWorkflowName(t workflow.WorkflowName) texit.WorkflowName {
	switch t {
	case workflow.WorkflowNameDeprovisionNode:
		return texit.WorkflowNameDeprovisionNode
	case workflow.WorkflowNameProvisionNode:
		return texit.WorkflowNameProvisionNode
	default:
		return texit.WorkflowNameUnknown
	}
}

func ExecutionToSummary(e *workflow.Execution) texit.ExecutionSummary {
	return texit.ExecutionSummary{
		Identifier: e.Identifier.String(),
		Workflow:   TranslateWorkflowName(e.Workflow),
		Status:     TranslateExecutionStatus(e.Status),
		StartedAt:  float64(e.Created.Unix()),
		EndedAt:    maybeMakeTime(e.Finished),
		Result:     maybeMakeString(e.Results.String()),
	}
}

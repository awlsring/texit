package conversion

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func TranslateWorkflowName(n workflow.WorkflowName) v1.WorkflowName {
	switch n {
	case workflow.WorkflowNameDeprovisionNode:
		return v1.WorkflowName_WORKFLOW_NAME_DEPROVISION_NODE
	case workflow.WorkflowNameProvisionNode:
		return v1.WorkflowName_WORKFLOW_NAME_PROVISION_NODE
	default:
		return v1.WorkflowName_WORKFLOW_NAME_UNKNOWN_UNSPECIFIED
	}
}

func TranslateExecutionStatus(s workflow.Status) v1.ExecutionStatus {
	switch s {
	case workflow.StatusPending:
		return v1.ExecutionStatus_EXECUTION_STATUS_PENDING
	case workflow.StatusRunning:
		return v1.ExecutionStatus_EXECUTION_STATUS_RUNNING
	case workflow.StatusComplete:
		return v1.ExecutionStatus_EXECUTION_STATUS_COMPLETED
	case workflow.StatusFailed:
		return v1.ExecutionStatus_EXECUTION_STATUS_FAILED
	default:
		return v1.ExecutionStatus_EXECUTION_STATUS_UNKNOWN_UNSPECIFIED
	}
}

func ExecutionToSummary(ex *workflow.Execution) *v1.ExecutionSummary {
	return &v1.ExecutionSummary{
		Id:           ex.Identifier.String(),
		WorkflowName: TranslateWorkflowName(ex.Workflow),
		Status:       TranslateExecutionStatus(ex.Status),
		CreatedAt:    ex.Created.Format(time.RFC3339Nano),
		UpdatedAt:    ex.Updated.Format(time.RFC3339Nano),
		FinishedAt:   ex.Finished.Format(time.RFC3339Nano),
	}
}

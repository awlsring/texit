package apiv1

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/texit/pkg/gen/client/v1"
)

func TranslateWorkflowName(n v1.WorkflowName) workflow.WorkflowName {
	switch n {
	case v1.WorkflowName_WORKFLOW_NAME_PROVISION_NODE:
		return workflow.WorkflowNameProvisionNode
	case v1.WorkflowName_WORKFLOW_NAME_DEPROVISION_NODE:
		return workflow.WorkflowNameDeprovisionNode
	default:
		return workflow.WorkflowNameUnknown
	}
}

func TranslateExecutionStatus(s v1.ExecutionStatus) workflow.Status {
	switch s {
	case v1.ExecutionStatus_EXECUTION_STATUS_PENDING:
		return workflow.StatusPending
	case v1.ExecutionStatus_EXECUTION_STATUS_RUNNING:
		return workflow.StatusRunning
	case v1.ExecutionStatus_EXECUTION_STATUS_COMPLETED:
		return workflow.StatusComplete
	case v1.ExecutionStatus_EXECUTION_STATUS_FAILED:
		return workflow.StatusFailed
	default:
		return workflow.StatusUnknown
	}
}

func getFinished(t string) (*time.Time, error) {
	if t == "" {
		return nil, nil
	}
	finished, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return nil, err
	}

	return &finished, nil
}

func SummaryToExecution(s *v1.ExecutionSummary) (*workflow.Execution, error) {
	id, err := workflow.ExecutionIdentifierFromString(s.GetId())
	if err != nil {
		return nil, err
	}

	created, err := time.Parse(time.RFC3339, s.GetCreatedAt())
	if err != nil {
		return nil, err
	}

	updated, err := time.Parse(time.RFC3339, s.GetUpdatedAt())
	if err != nil {
		return nil, err
	}

	finished, err := getFinished(s.GetFinishedAt())
	if err != nil {
		return nil, err
	}

	return &workflow.Execution{
		Identifier: id,
		Workflow:   TranslateWorkflowName(s.WorkflowName),
		Status:     TranslateExecutionStatus(s.Status),
		Created:    created,
		Updated:    updated,
		Finished:   finished,
	}, nil
}

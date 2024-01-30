package conversion

import (
	"testing"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"github.com/stretchr/testify/assert"
)

func TestTranslateWorkflowName(t *testing.T) {
	assert.Equal(t, v1.WorkflowName_WORKFLOW_NAME_DEPROVISION_NODE, TranslateWorkflowName(workflow.WorkflowNameDeprovisionNode))
	assert.Equal(t, v1.WorkflowName_WORKFLOW_NAME_PROVISION_NODE, TranslateWorkflowName(workflow.WorkflowNameProvisionNode))
	assert.Equal(t, v1.WorkflowName_WORKFLOW_NAME_UNKNOWN_UNSPECIFIED, TranslateWorkflowName(workflow.WorkflowNameUnknown))
}

func TestTranslateExecutionStatus(t *testing.T) {
	assert.Equal(t, v1.ExecutionStatus_EXECUTION_STATUS_PENDING, TranslateExecutionStatus(workflow.StatusPending))
	assert.Equal(t, v1.ExecutionStatus_EXECUTION_STATUS_RUNNING, TranslateExecutionStatus(workflow.StatusRunning))
	assert.Equal(t, v1.ExecutionStatus_EXECUTION_STATUS_COMPLETED, TranslateExecutionStatus(workflow.StatusComplete))
	assert.Equal(t, v1.ExecutionStatus_EXECUTION_STATUS_FAILED, TranslateExecutionStatus(workflow.StatusFailed))
	assert.Equal(t, v1.ExecutionStatus_EXECUTION_STATUS_UNKNOWN_UNSPECIFIED, TranslateExecutionStatus(workflow.StatusUnknown))
}

func TestExecutionToSummary(t *testing.T) {
	exId, err := workflow.ExecutionIdentifierFromString("test-id")
	assert.NoError(t, err)

	now := time.Now()

	ex := &workflow.Execution{
		Identifier: exId,
		Workflow:   workflow.WorkflowNameProvisionNode,
		Status:     workflow.StatusRunning,
		Created:    now,
		Updated:    now,
		Finished:   &now,
	}

	summary := ExecutionToSummary(ex)

	assert.Equal(t, ex.Identifier.String(), summary.Id)
	assert.Equal(t, TranslateWorkflowName(ex.Workflow), summary.WorkflowName)
	assert.Equal(t, TranslateExecutionStatus(ex.Status), summary.Status)
	assert.Equal(t, ex.Created.Format(time.RFC3339Nano), summary.CreatedAt)
	assert.Equal(t, ex.Updated.Format(time.RFC3339Nano), summary.UpdatedAt)
	assert.Equal(t, ex.Finished.Format(time.RFC3339Nano), summary.FinishedAt)
}

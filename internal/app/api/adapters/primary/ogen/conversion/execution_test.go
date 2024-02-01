package conversion

import (
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestTranslateExecutionStatus(t *testing.T) {
	assert.Equal(t, texit.ExecutionStatusPending, TranslateExecutionStatus(workflow.StatusPending))
	assert.Equal(t, texit.ExecutionStatusRunning, TranslateExecutionStatus(workflow.StatusRunning))
	assert.Equal(t, texit.ExecutionStatusCompleted, TranslateExecutionStatus(workflow.StatusComplete))
	assert.Equal(t, texit.ExecutionStatusFailed, TranslateExecutionStatus(workflow.StatusFailed))
	assert.Equal(t, texit.ExecutionStatusUnknown, TranslateExecutionStatus(workflow.Status(999))) // Unknown status
}

func TestTranslateWorkflowName(t *testing.T) {
	assert.Equal(t, texit.WorkflowNameDeprovisionNode, TranslateWorkflowName(workflow.WorkflowNameDeprovisionNode))
	assert.Equal(t, texit.WorkflowNameProvisionNode, TranslateWorkflowName(workflow.WorkflowNameProvisionNode))
	assert.Equal(t, texit.WorkflowNameUnknown, TranslateWorkflowName(workflow.WorkflowNameUnknown)) // Unknown workflow name
}

func TestExecutionToSummary(t *testing.T) {
	execution := &workflow.Execution{
		Identifier: workflow.ExecutionIdentifier("test-id"),
		Workflow:   workflow.WorkflowNameProvisionNode,
		Status:     workflow.StatusRunning,
		Created:    time.Now(),
		Finished:   nil,
	}

	summary := ExecutionToSummary(execution)

	assert.Equal(t, execution.Identifier.String(), summary.Identifier)
	assert.Equal(t, TranslateWorkflowName(execution.Workflow), summary.Workflow)
	assert.Equal(t, TranslateExecutionStatus(execution.Status), summary.Status)
	assert.Equal(t, float64(execution.Created.Unix()), summary.StartedAt)
	assert.Equal(t, summary.EndedAt, texit.OptFloat64{})
	assert.Equal(t, maybeMakeString(""), summary.Result)
}

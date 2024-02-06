package sqlite_execution_repository

import (
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToExecution(t *testing.T) {
	now := time.Now()

	t.Run("returns Execution with same values as ExecutionSqlRecord", func(t *testing.T) {
		record := &ExecutionSqlRecord{
			Identifier: uuid.New().String(),
			Workflow:   workflow.WorkflowNameProvisionNode.String(),
			Status:     workflow.StatusRunning.String(),
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		result := record.ToExecution()
		assert.Equal(t, record.Identifier, result.Identifier.String())
		assert.Equal(t, record.Workflow, result.Workflow.String())
		assert.Equal(t, record.Status, result.Status.String())
		assert.Equal(t, record.CreatedAt, result.Created)
		assert.Equal(t, record.UpdatedAt, result.Updated)
		assert.Nil(t, result.Finished)
		assert.Equal(t, result.Results.String(), "")
	})

	t.Run("returns optional fields when set", func(t *testing.T) {
		results := "{\"nodeId\":\"node\",\"failedStep\":\"step\",\"errors\":[\"error\"]}"

		record := &ExecutionSqlRecord{
			Identifier: uuid.New().String(),
			Workflow:   workflow.WorkflowNameProvisionNode.String(),
			Status:     workflow.StatusRunning.String(),
			CreatedAt:  now,
			UpdatedAt:  now,
			FinishedAt: &now,
			Results:    &results,
		}
		result := record.ToExecution()
		assert.Equal(t, record.Identifier, result.Identifier.String())
		assert.Equal(t, record.Workflow, result.Workflow.String())
		assert.Equal(t, record.Status, result.Status.String())
		assert.Equal(t, record.CreatedAt, result.Created)
		assert.Equal(t, record.UpdatedAt, result.Updated)
		assert.Equal(t, *record.FinishedAt, *result.Finished)
		assert.Equal(t, result.Results.String(), results)
	})

}

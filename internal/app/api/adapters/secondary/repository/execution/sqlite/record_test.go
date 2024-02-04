package sqlite_execution_repository

import (
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestResultsToStringList(t *testing.T) {
	t.Run("returns nil when input is nil", func(t *testing.T) {
		var input *string
		expected := []string(nil)
		result := resultsToStringList(input)
		assert.Equal(t, expected, result)
	})

	t.Run("returns nil when input is an empty string", func(t *testing.T) {
		input := ""
		expected := []string(nil)
		result := resultsToStringList(&input)
		assert.Equal(t, expected, result)
	})

	t.Run("returns a list of strings when input is a comma-separated string", func(t *testing.T) {
		input := "test1,test2,test3"
		expected := []string{"test1", "test2", "test3"}
		result := resultsToStringList(&input)
		assert.Equal(t, expected, result)
	})

	t.Run("returns a list with one string when input is a string without commas", func(t *testing.T) {
		input := "test1"
		expected := []string{"test1"}
		result := resultsToStringList(&input)
		assert.Equal(t, expected, result)
	})
}

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
		assert.Nil(t, result.Results)
	})

	t.Run("returns optional fields when set", func(t *testing.T) {
		results := "test1,test2,test3"
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
		assert.Equal(t, []string{"test1", "test2", "test3"}, result.Results)
	})

}

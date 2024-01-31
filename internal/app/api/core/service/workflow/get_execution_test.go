package workflow

import (
	"context"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/stretchr/testify/assert"
)

func TestGetExecution(t *testing.T) {
	ctx := context.Background()
	id := workflow.ExecutionIdentifier("test-execution")

	s := &Service{
		executions: map[string]*workflow.Execution{
			id.String(): {
				Identifier: id,
				Status:     workflow.StatusComplete,
			},
		},
	}

	exec, err := s.GetExecution(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, exec)
	assert.Equal(t, id, exec.Identifier)
	assert.Equal(t, workflow.StatusComplete, exec.Status)
}

func TestGetExecutionNotFound(t *testing.T) {
	ctx := context.Background()
	id := workflow.ExecutionIdentifier("test-execution")

	s := &Service{
		executions: map[string]*workflow.Execution{},
	}

	exec, err := s.GetExecution(ctx, id)

	assert.Error(t, err)
	assert.Nil(t, exec)
}

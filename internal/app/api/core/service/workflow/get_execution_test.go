package workflow

import (
	"context"
	"testing"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetExecution(t *testing.T) {
	ctx := context.Background()
	id := workflow.ExecutionIdentifier("test-execution")

	mockRepo := mocks.NewMockExecution_repository(t)

	s := &Service{
		excRepo: mockRepo,
	}

	mockRepo.EXPECT().GetExecution(ctx, id).Return(&workflow.Execution{
		Identifier: id,
		Status:     workflow.StatusComplete,
	}, nil)

	exec, err := s.GetExecution(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, exec)
	assert.Equal(t, id, exec.Identifier)
	assert.Equal(t, workflow.StatusComplete, exec.Status)
}

func TestGetExecutionNotFound(t *testing.T) {
	ctx := context.Background()
	id := workflow.ExecutionIdentifier("test-execution")

	mockRepo := mocks.NewMockExecution_repository(t)

	s := &Service{
		excRepo: mockRepo,
	}

	mockRepo.EXPECT().GetExecution(ctx, id).Return(nil, repository.ErrExecutionNotFound)

	exec, err := s.GetExecution(ctx, id)

	assert.Error(t, err)
	assert.Nil(t, exec)
}

package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestGetExecution(t *testing.T) {
	ctx := context.Background()
	id := workflow.ExecutionIdentifier("test-execution")

	req := texit.GetExecutionParams{
		Identifier: id.String(),
	}

	t.Run("returns execution when get execution is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().GetExecution(ctx, req).Return(&texit.GetExecutionResponseContent{
			Summary: texit.ExecutionSummary{
				Identifier: id.String(),
				Status:     texit.ExecutionStatusCompleted,
			},
		}, nil)

		exec, err := g.GetExecution(ctx, id)

		assert.NoError(t, err)
		assert.NotNil(t, exec)
		assert.Equal(t, id, exec.Identifier)
		assert.Equal(t, workflow.StatusComplete, exec.Status)
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when get execution fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().GetExecution(ctx, req).Return(nil, errors.New("get execution failed"))

		exec, err := g.GetExecution(ctx, id)

		assert.Error(t, err)
		assert.EqualError(t, err, "get execution failed")
		assert.Nil(t, exec)
		mockClient.AssertExpectations(t)
	})
}

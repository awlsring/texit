package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestDeprovisionNode(t *testing.T) {
	ctx := context.Background()
	id := node.Identifier("test-node")

	req := texit.DeprovisionNodeParams{
		Identifier: id.String(),
	}

	t.Run("returns execution id when deprovision is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		apiGateway := New(mockClient)
		mockClient.EXPECT().DeprovisionNode(ctx, req).Return(&texit.DeprovisionNodeResponseContent{Execution: "test-execution"}, nil)

		exId, err := apiGateway.DeprovisionNode(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, workflow.ExecutionIdentifier("test-execution"), exId)
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when deprovision fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		apiGateway := New(mockClient)
		mockClient.EXPECT().DeprovisionNode(ctx, req).Return(nil, errors.New("deprovision failed"))

		exId, err := apiGateway.DeprovisionNode(ctx, id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, gateway.ErrInternalServerError)
		assert.Equal(t, workflow.ExecutionIdentifier(""), exId)
		mockClient.AssertExpectations(t)
	})
}

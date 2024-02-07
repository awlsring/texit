package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestApiGateway_GetNodeStatus(t *testing.T) {
	ctx := context.Background()
	id := node.Identifier("test-node")

	req := texit.GetNodeStatusParams{
		Identifier: id.String(),
	}

	t.Run("returns status when get node status is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().GetNodeStatus(ctx, req).Return(&texit.GetNodeStatusResponseContent{
			Status: texit.NodeStatusRunning,
		}, nil)

		status, err := g.GetNodeStatus(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, node.StatusRunning, status)
	})

	t.Run("returns unknown status when get node status fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().GetNodeStatus(ctx, req).Return(nil, errors.New("get node status failed"))

		status, err := g.GetNodeStatus(ctx, id)

		assert.Error(t, err)
		assert.ErrorIs(t, err, gateway.ErrInternalServerError)
		assert.Equal(t, node.StatusUnknown, status)
	})

	t.Run("returns translated status when get node status returns non-running status", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().GetNodeStatus(ctx, req).Return(&texit.GetNodeStatusResponseContent{
			Status: texit.NodeStatusStopped,
		}, nil)

		status, err := g.GetNodeStatus(ctx, id)

		assert.NoError(t, err)
		assert.Equal(t, node.StatusStopped, status)
	})
}

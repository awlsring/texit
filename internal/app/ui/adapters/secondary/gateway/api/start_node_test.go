package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestApiGateway_StartNode(t *testing.T) {
	ctx := context.Background()

	nodeId, _ := node.IdentifierFromString("test-node")

	t.Run("returns no error when start node is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().StartNode(ctx, texit.StartNodeParams{
			Identifier: nodeId.String(),
		}).Return(&texit.StartNodeResponseContent{Success: true}, nil)

		err := g.StartNode(ctx, nodeId)

		assert.NoError(t, err)
	})

	t.Run("returns error when start node fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().StartNode(ctx, texit.StartNodeParams{
			Identifier: nodeId.String(),
		}).Return(nil, errors.New("start node failed"))

		err := g.StartNode(ctx, nodeId)

		assert.Error(t, err)
		assert.EqualError(t, err, "start node failed")
	})
}

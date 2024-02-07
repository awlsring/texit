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

func TestApiGateway_StopNode(t *testing.T) {
	ctx := context.Background()

	nodeId, _ := node.IdentifierFromString("test-node")

	t.Run("returns no error when stop node is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}

		mockClient.EXPECT().StopNode(ctx, texit.StopNodeParams{
			Identifier: nodeId.String(),
		}).Return(&texit.StopNodeResponseContent{Success: true}, nil)

		err := g.StopNode(ctx, nodeId)

		assert.NoError(t, err)
	})

	t.Run("returns error when stop node fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}

		mockClient.EXPECT().StopNode(ctx, texit.StopNodeParams{
			Identifier: nodeId.String(),
		}).Return(nil, errors.New("stop node failed"))

		err := g.StopNode(ctx, nodeId)

		assert.Error(t, err)
		assert.ErrorIs(t, err, gateway.ErrInternalServerError)
	})
}

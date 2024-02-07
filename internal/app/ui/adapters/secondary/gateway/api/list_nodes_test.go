package api_gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestApiGateway_ListNodes(t *testing.T) {
	ctx := context.Background()

	nowFloat := float64(time.Now().Unix())
	testSummaries := []texit.NodeSummary{
		{
			Identifier:              "test-node-1",
			Provider:                "test-provider",
			Location:                "test-region",
			ProviderNodeIdentifier:  "test-provider-node",
			Tailnet:                 "test-tailnet",
			TailnetDeviceName:       "test-tailnet-device",
			TailnetDeviceIdentifier: "test-tailnet-device-id",
			Ephemeral:               false,
			Created:                 nowFloat,
			Updated:                 nowFloat,
		},
		{
			Identifier:              "test-node-2",
			Provider:                "test-provider",
			Location:                "test-region",
			ProviderNodeIdentifier:  "test-provider-node",
			Tailnet:                 "test-tailnet",
			TailnetDeviceName:       "test-tailnet-device",
			TailnetDeviceIdentifier: "test-tailnet-device-id",
			Ephemeral:               false,
			Created:                 nowFloat,
			Updated:                 nowFloat,
		},
	}

	t.Run("returns nodes when list nodes is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListNodes(ctx).Return(&texit.ListNodesResponseContent{
			Summaries: testSummaries,
		}, nil)

		nodes, err := g.ListNodes(ctx)

		assert.NoError(t, err)
		assert.Len(t, nodes, len(testSummaries))
		for i, node := range nodes {
			assert.Equal(t, testSummaries[i].Identifier, node.Identifier.String())
		}
	})

	t.Run("returns error when list nodes fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListNodes(ctx).Return(nil, errors.New("list nodes failed"))

		nodes, err := g.ListNodes(ctx)

		assert.Error(t, err)
		assert.ErrorIs(t, err, gateway.ErrInternalServerError)
		assert.Nil(t, nodes)
	})

	t.Run("returns error when summary to node conversion fails", func(t *testing.T) {
		invalidSummary := texit.NodeSummary{
			Identifier: "",
		}
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListNodes(ctx).Return(&texit.ListNodesResponseContent{
			Summaries: []texit.NodeSummary{invalidSummary},
		}, nil)

		nodes, err := g.ListNodes(ctx)

		assert.Error(t, err)
		assert.Nil(t, nodes)
	})
}

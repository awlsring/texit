package api_gateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestGetNode(t *testing.T) {
	ctx := context.Background()
	id := node.Identifier("test-node")

	req := texit.DescribeNodeParams{
		Identifier: id.String(),
	}

	nowFloat := float64(time.Now().Unix())

	sum := texit.NodeSummary{
		Identifier:              id.String(),
		Provider:                "test-provider",
		Location:                "test-region",
		ProviderNodeIdentifier:  "test-provider-node",
		Tailnet:                 "test-tailnet",
		TailnetDeviceName:       "test-tailnet-device",
		TailnetDeviceIdentifier: "test-tailnet-device-id",
		Ephemeral:               false,
		Created:                 nowFloat,
		Updated:                 nowFloat,
	}

	t.Run("returns node when get node is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := New(mockClient)
		mockClient.EXPECT().DescribeNode(ctx, req).Return(&texit.DescribeNodeResponseContent{
			Summary: sum,
		}, nil)

		n, err := g.DescribeNode(ctx, id)

		assert.NoError(t, err)
		assert.NotNil(t, n)
		assert.Equal(t, id, n.Identifier)
	})

	t.Run("returns error when get node fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := New(mockClient)
		mockClient.EXPECT().DescribeNode(ctx, req).Return(nil, errors.New("get node failed"))

		n, err := g.DescribeNode(ctx, id)

		assert.Error(t, err)
		assert.EqualError(t, err, "get node failed")
		assert.Nil(t, n)
	})

	t.Run("returns error when summary to node conversion fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := New(mockClient)
		mockClient.EXPECT().DescribeNode(ctx, req).Return(&texit.DescribeNodeResponseContent{
			Summary: texit.NodeSummary{
				Identifier: id.String(),
			},
		}, nil)

		n, err := g.DescribeNode(ctx, id)

		assert.Error(t, err)
		assert.Nil(t, n)
	})
}

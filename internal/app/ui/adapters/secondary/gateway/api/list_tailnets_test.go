package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestApiGateway_ListTailnets(t *testing.T) {
	ctx := context.Background()

	testSummaries := []texit.TailnetSummary{
		{
			Name: "test-tailnet-1",
			Type: texit.TailnetTypeTailscale,
		},
		{
			Name: "test-tailnet-2",
			Type: texit.TailnetTypeTailscale,
		},
	}

	t.Run("returns tailnets when list tailnets is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListTailnets(ctx).Return(&texit.ListTailnetsResponseContent{
			Summaries: testSummaries,
		}, nil)

		tailnets, err := g.ListTailnets(ctx)

		assert.NoError(t, err)
		assert.Len(t, tailnets, len(testSummaries))
		for i, tailnet := range tailnets {
			assert.Equal(t, testSummaries[i].Name, tailnet.Name.String())
		}
	})

	t.Run("returns error when list tailnets fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListTailnets(ctx).Return(nil, errors.New("list tailnets failed"))

		tailnets, err := g.ListTailnets(ctx)

		assert.Error(t, err)
		assert.EqualError(t, err, "list tailnets failed")
		assert.Nil(t, tailnets)
	})

	t.Run("returns error when summary to tailnet conversion fails", func(t *testing.T) {
		invalidSummary := texit.TailnetSummary{
			Name: "",
		}
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListTailnets(ctx).Return(&texit.ListTailnetsResponseContent{
			Summaries: []texit.TailnetSummary{invalidSummary},
		}, nil)

		tailnets, err := g.ListTailnets(ctx)

		assert.Error(t, err)
		assert.Nil(t, tailnets)
	})
}

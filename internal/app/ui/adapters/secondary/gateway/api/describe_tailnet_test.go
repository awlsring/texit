package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestDescribeTailnet(t *testing.T) {
	ctx := context.Background()
	identifier, _ := tailnet.IdentifierFromString("test-tailnet")

	req := texit.DescribeTailnetParams{
		Name: identifier.String(),
	}

	t.Run("returns tailnet when describe tailnet is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().DescribeTailnet(ctx, req).Return(&texit.DescribeTailnetResponseContent{
			Summary: texit.TailnetSummary{
				Name: identifier.String(),
				Type: texit.TailnetTypeHeadscale,
			},
		}, nil)

		tail, err := g.DescribeTailnet(ctx, identifier)

		assert.NoError(t, err)
		assert.NotNil(t, tail)
		assert.Equal(t, identifier, tail.Name)
	})

	t.Run("returns error when describe tailnet fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().DescribeTailnet(ctx, req).Return(nil, errors.New("describe tailnet failed"))

		tail, err := g.DescribeTailnet(ctx, identifier)

		assert.Error(t, err)
		assert.EqualError(t, err, "describe tailnet failed")
		assert.Nil(t, tail)
	})

	t.Run("returns error when summary to tailnet conversion fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().DescribeTailnet(ctx, req).Return(&texit.DescribeTailnetResponseContent{
			Summary: texit.TailnetSummary{
				Name: "",
			},
		}, nil)

		tail, err := g.DescribeTailnet(ctx, identifier)

		assert.Error(t, err)
		assert.Nil(t, tail)
	})
}

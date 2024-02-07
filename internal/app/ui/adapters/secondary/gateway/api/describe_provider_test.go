package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestGetProvider(t *testing.T) {
	ctx := context.Background()
	identifier, _ := provider.IdentifierFromString("test-provider")

	req := texit.DescribeProviderParams{
		Name: identifier.String(),
	}

	t.Run("returns provider when get provider is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().DescribeProvider(ctx, req).Return(&texit.DescribeProviderResponseContent{
			Summary: texit.ProviderSummary{
				Name: identifier.String(),
			},
		}, nil)

		prov, err := g.DescribeProvider(ctx, identifier)

		assert.NoError(t, err)
		assert.NotNil(t, prov)
		assert.Equal(t, identifier, prov.Name)
	})

	t.Run("returns error when get provider fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().DescribeProvider(ctx, req).Return(nil, errors.New("get provider failed"))

		prov, err := g.DescribeProvider(ctx, identifier)

		assert.Error(t, err)
		assert.ErrorIs(t, err, gateway.ErrInternalServerError)
		assert.Nil(t, prov)
	})

	t.Run("returns error when summary to provider conversion fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().DescribeProvider(ctx, req).Return(&texit.DescribeProviderResponseContent{
			Summary: texit.ProviderSummary{
				Name: "",
			},
		}, nil)

		prov, err := g.DescribeProvider(ctx, identifier)

		assert.Error(t, err)
		assert.Nil(t, prov)
	})
}

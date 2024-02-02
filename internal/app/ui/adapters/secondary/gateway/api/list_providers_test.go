package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestApiGateway_ListProviders(t *testing.T) {
	ctx := context.Background()

	testSummaries := []texit.ProviderSummary{
		{
			Name: "test-provider-1",
			Type: texit.ProviderTypeAWSEcs,
		},
		{
			Name: "test-provider-2",
			Type: texit.ProviderTypeAWSEcs,
		},
	}

	t.Run("returns providers when list providers is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListProviders(ctx).Return(&texit.ListProvidersResponseContent{
			Summaries: testSummaries,
		}, nil)

		providers, err := g.ListProviders(ctx)

		assert.NoError(t, err)
		assert.Len(t, providers, len(testSummaries))
		for i, provider := range providers {
			assert.Equal(t, testSummaries[i].Name, provider.Name.String())
		}
	})

	t.Run("returns error when list providers fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListProviders(ctx).Return(nil, errors.New("list providers failed"))

		providers, err := g.ListProviders(ctx)

		assert.Error(t, err)
		assert.EqualError(t, err, "list providers failed")
		assert.Nil(t, providers)
	})

	t.Run("returns error when summary to provider conversion fails", func(t *testing.T) {
		invalidSummary := texit.ProviderSummary{
			Name: "",
		}
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ListProviders(ctx).Return(&texit.ListProvidersResponseContent{
			Summaries: []texit.ProviderSummary{invalidSummary},
		}, nil)

		providers, err := g.ListProviders(ctx)

		assert.Error(t, err)
		assert.Nil(t, providers)
	})
}

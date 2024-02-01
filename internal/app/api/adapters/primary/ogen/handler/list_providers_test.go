package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListProviders(t *testing.T) {
	ctx := context.Background()

	mockProviderSvc := mocks.NewMockProvider_service(t)
	h := New(nil, nil, mockProviderSvc, nil)

	testProviders := []*provider.Provider{
		{
			Name: provider.Identifier("test-provider-1"),
		},
		{
			Name: provider.Identifier("test-provider-2"),
		},
	}

	mockProviderSvc.EXPECT().List(ctx).Return(testProviders, nil)

	res, err := h.ListProviders(ctx)

	assert.NoError(t, err)
	assert.Len(t, res.Summaries, len(testProviders))
}

func TestListProvidersError(t *testing.T) {
	ctx := context.Background()

	mockProviderSvc := mocks.NewMockProvider_service(t)
	h := New(nil, nil, mockProviderSvc, nil)

	mockProviderSvc.EXPECT().List(ctx).Return(nil, errors.New("test error"))

	res, err := h.ListProviders(ctx)

	assert.Error(t, err)
	assert.Nil(t, res)
}

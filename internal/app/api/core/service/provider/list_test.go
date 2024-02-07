package provider

import (
	"context"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/stretchr/testify/assert"
)

func TestListProviders(t *testing.T) {
	ctx := context.Background()

	mockProvider1 := &provider.Provider{
		Name: provider.Identifier("test-provider-1"),
	}

	mockProvider2 := &provider.Provider{
		Name: provider.Identifier("test-provider-2"),
	}

	s := &Service{
		provMap: map[provider.Identifier]*provider.Provider{
			mockProvider1.Name: mockProvider1,
			mockProvider2.Name: mockProvider2,
		},
	}

	result, err := s.List(ctx)

	assert.NoError(t, err)
	assert.Contains(t, result, mockProvider1)
	assert.Contains(t, result, mockProvider2)
}

func TestListProvidersEmpty(t *testing.T) {
	ctx := context.Background()

	s := &Service{
		provMap: map[provider.Identifier]*provider.Provider{},
	}

	result, err := s.List(ctx)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

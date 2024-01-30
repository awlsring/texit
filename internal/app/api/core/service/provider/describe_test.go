package provider

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDescribe(t *testing.T) {
	ctx := context.Background()

	mockProvider := &provider.Provider{
		Name: provider.Identifier("test-provider"),
	}

	s := &Service{
		provMap: map[provider.Identifier]*provider.Provider{
			mockProvider.Name: mockProvider,
		},
	}

	result, err := s.Describe(ctx, mockProvider.Name)

	assert.NoError(t, err)
	assert.Equal(t, mockProvider, result)
}

func TestDescribeUnknownProvider(t *testing.T) {
	ctx := context.Background()

	s := &Service{
		provMap: map[provider.Identifier]*provider.Provider{},
	}

	result, err := s.Describe(ctx, provider.Identifier("unknown-provider"))

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, service.ErrUnknownProvider))
}

package tailnet

import (
	"context"
	"testing"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestDescribeTailnet(t *testing.T) {
	ctx := context.Background()

	mockTailnet := &tailnet.Tailnet{
		Name: tailnet.Identifier("test-tailnet"),
	}

	s := &Service{
		tailnetMap: map[tailnet.Identifier]*tailnet.Tailnet{
			mockTailnet.Name: mockTailnet,
		},
	}

	result, err := s.Describe(ctx, mockTailnet.Name)

	assert.NoError(t, err)
	assert.Equal(t, mockTailnet, result)
}

func TestDescribeUnknownTailnet(t *testing.T) {
	ctx := context.Background()

	s := &Service{
		tailnetMap: map[tailnet.Identifier]*tailnet.Tailnet{},
	}

	result, err := s.Describe(ctx, tailnet.Identifier("unknown-tailnet"))

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.True(t, errors.Is(err, service.ErrUnknownProvider))
}

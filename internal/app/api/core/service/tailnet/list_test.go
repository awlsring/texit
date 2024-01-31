package tailnet

import (
	"context"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/stretchr/testify/assert"
)

func TestListTailnets(t *testing.T) {
	ctx := context.Background()

	mockTailnet1 := &tailnet.Tailnet{
		Name: tailnet.Identifier("test-tailnet-1"),
	}

	mockTailnet2 := &tailnet.Tailnet{
		Name: tailnet.Identifier("test-tailnet-2"),
	}

	s := &Service{
		tailnetMap: map[tailnet.Identifier]*tailnet.Tailnet{
			mockTailnet1.Name: mockTailnet1,
			mockTailnet2.Name: mockTailnet2,
		},
	}

	result, err := s.List(ctx)

	assert.NoError(t, err)
	assert.Equal(t, []*tailnet.Tailnet{mockTailnet1, mockTailnet2}, result)
}

func TestListTailnetsEmpty(t *testing.T) {
	ctx := context.Background()

	s := &Service{
		tailnetMap: map[tailnet.Identifier]*tailnet.Tailnet{},
	}

	result, err := s.List(ctx)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

package tailscale_gateway

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func TestCreatePreauthKey(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
		user:   testUser,
	}

	testKey := "test-key"
	mockResponse := tailscale.Key{
		Key: testKey,
	}

	mockClient.EXPECT().CreateKey(mock.Anything, mock.Anything).Return(mockResponse, nil)

	resultKey, err := g.CreatePreauthKey(ctx, true)

	assert.NoError(t, err)
	assert.Equal(t, tailnet.PreauthKey(testKey), resultKey)
}

func TestCreatePreauthKey_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	mockClient.EXPECT().CreateKey(mock.Anything, mock.Anything).Return(tailscale.Key{}, errors.New("test error"))

	_, err := g.CreatePreauthKey(ctx, true)

	assert.Error(t, err)
}

func TestCreatePreauthKey_Ephemeral(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testKey := "test-key"
	mockResponse := tailscale.Key{
		Key: testKey,
	}

	mockClient.EXPECT().CreateKey(mock.Anything, mock.Anything).Return(mockResponse, nil)

	resultKey, err := g.CreatePreauthKey(ctx, false)

	assert.NoError(t, err)
	assert.Equal(t, tailnet.PreauthKey(testKey), resultKey)
}

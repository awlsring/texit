package tailscale_gateway

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeletePreauthKey(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testKey := tailnet.PreauthKey("test-key")
	mockClient.EXPECT().DeleteKey(mock.Anything, testKey.String()).Return(nil)

	err := g.DeletePreauthKey(ctx, testKey)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDeletePreauthKey_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testKey := tailnet.PreauthKey("test-key")
	mockClient.EXPECT().DeleteKey(mock.Anything, testKey.String()).Return(errors.New("test error"))

	err := g.DeletePreauthKey(ctx, testKey)

	assert.Error(t, err)
	mockClient.AssertExpectations(t)
}

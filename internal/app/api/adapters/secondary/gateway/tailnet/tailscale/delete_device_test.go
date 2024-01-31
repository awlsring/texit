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

func TestDeleteDevice(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testID := tailnet.DeviceIdentifier("test-id")
	mockClient.EXPECT().DeleteDevice(mock.Anything, testID.String()).Return(nil)

	err := g.DeleteDevice(ctx, testID)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestDeleteDevice_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testID := tailnet.DeviceIdentifier("test-id")
	mockClient.EXPECT().DeleteDevice(mock.Anything, testID.String()).Return(errors.New("test error"))

	err := g.DeleteDevice(ctx, testID)

	assert.Error(t, err)
	mockClient.AssertExpectations(t)
}

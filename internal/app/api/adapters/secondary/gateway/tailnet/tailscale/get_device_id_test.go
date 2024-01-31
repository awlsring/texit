package tailscale_gateway

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func TestGetDeviceId(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testName := tailnet.DeviceName("test-name")
	testID := "test-id"
	mockDevices := []tailscale.Device{
		{
			ID:       testID,
			Hostname: testName.String(),
		},
	}
	mockClient.EXPECT().Devices(mock.Anything).Return(mockDevices, nil)

	resultID, err := g.GetDeviceId(ctx, testName)

	assert.NoError(t, err)
	assert.Equal(t, tailnet.DeviceIdentifier(testID), resultID)
	mockClient.AssertExpectations(t)
}

func TestGetDeviceId_ListDevicesError(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testName := tailnet.DeviceName("test-name")
	mockClient.EXPECT().Devices(mock.Anything).Return(nil, errors.New("test error"))

	_, err := g.GetDeviceId(ctx, testName)

	assert.Error(t, err)
	mockClient.AssertExpectations(t)
}

func TestGetDeviceId_DeviceNotFound(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	g := &TailscaleGateway{
		client: mockClient,
	}

	testName := tailnet.DeviceName("test-name")
	mockDevices := []tailscale.Device{
		{
			ID:       "other-id",
			Hostname: "other-name",
		},
	}
	mockClient.EXPECT().Devices(mock.Anything).Return(mockDevices, nil)

	_, err := g.GetDeviceId(ctx, testName)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, gateway.ErrUnknownDevice))
	mockClient.AssertExpectations(t)
}

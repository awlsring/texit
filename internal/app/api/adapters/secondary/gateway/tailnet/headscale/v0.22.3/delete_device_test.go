package headscale_v0_22_3_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/client/headscale_service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDeleteDevice(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		user:   "test-user",
		client: mockClient,
	}

	testID := tailnet.DeviceIdentifier("test-id")
	mockClient.EXPECT().HeadscaleServiceDeleteMachine(mock.Anything).Return(&headscale_service.HeadscaleServiceDeleteMachineOK{}, nil)

	err := g.DeleteDevice(ctx, testID)

	assert.NoError(t, err)
}

func TestDeleteDevice_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testID := tailnet.DeviceIdentifier("test-id")
	mockClient.EXPECT().HeadscaleServiceDeleteMachine(mock.Anything).Return(nil, errors.New("test error"))

	err := g.DeleteDevice(ctx, testID)

	assert.Error(t, err)
}

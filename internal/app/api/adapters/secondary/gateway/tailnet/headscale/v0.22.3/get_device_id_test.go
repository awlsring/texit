package headscale_v0_22_3_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/client/headscale_service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDeviceId(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testName := tailnet.DeviceName("test-name")
	testID := "test-id"
	mockResponse := &headscale_service.HeadscaleServiceListMachinesOK{
		Payload: &models.V1ListMachinesResponse{
			Machines: []*models.V1Machine{
				{
					Name: testName.String(),
					ID:   testID,
				},
			},
		},
	}

	mockClient.On("HeadscaleServiceListMachines", mock.Anything).Return(mockResponse, nil)

	resultID, err := g.GetDeviceId(ctx, testName)

	assert.NoError(t, err)
	assert.Equal(t, tailnet.DeviceIdentifier(testID), resultID)
}

func TestGetDeviceId_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testName := tailnet.DeviceName("test-name")
	mockClient.On("HeadscaleServiceListMachines", mock.Anything).Return(nil, errors.New("test error"))

	_, err := g.GetDeviceId(ctx, testName)

	assert.Error(t, err)
}

func TestGetDeviceId_NotFound(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
	}

	testName := tailnet.DeviceName("test-name")
	mockResponse := &headscale_service.HeadscaleServiceListMachinesOK{
		Payload: &models.V1ListMachinesResponse{
			Machines: []*models.V1Machine{},
		},
	}

	mockClient.On("HeadscaleServiceListMachines", mock.Anything).Return(mockResponse, nil)

	_, err := g.GetDeviceId(ctx, testName)

	assert.Error(t, err)
}

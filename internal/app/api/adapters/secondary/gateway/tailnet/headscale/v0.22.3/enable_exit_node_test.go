package headscale_v0_22_3_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/client/headscale_service"
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetRoutesForDevice(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testId := tailnet.DeviceIdentifier("1")
	mockResponse := &headscale_service.HeadscaleServiceGetRoutesOK{
		Payload: &models.V1GetRoutesResponse{
			Routes: []*models.V1Route{
				{
					ID: "1",
					Machine: &models.V1Machine{
						ID:   "1",
						Name: "test-id",
					},
				},
				{
					ID: "2",
					Machine: &models.V1Machine{
						ID:   "2",
						Name: "test-other-id",
					},
				},
			},
		},
	}

	mockClient.On("HeadscaleServiceGetRoutes", mock.Anything).Return(mockResponse, nil)

	resultRoutes, err := g.getRoutesForDevice(ctx, testId)

	assert.NoError(t, err)
	assert.Equal(t, []string{"1"}, resultRoutes)
}

func TestGetRoutesForDeviceNoRoutes(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testId := tailnet.DeviceIdentifier("1")
	mockResponse := &headscale_service.HeadscaleServiceGetRoutesOK{
		Payload: &models.V1GetRoutesResponse{
			Routes: []*models.V1Route{
				{
					ID: "1",
					Machine: &models.V1Machine{
						ID:   "2",
						Name: "test-id",
					},
				},
				{
					ID: "2",
					Machine: &models.V1Machine{
						ID:   "2",
						Name: "test-other-id",
					},
				},
			},
		},
	}

	mockClient.On("HeadscaleServiceGetRoutes", mock.Anything).Return(mockResponse, nil)

	_, err := g.getRoutesForDevice(ctx, testId)

	assert.Error(t, err)
}

func TestGetRoutesForDevice_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testID := tailnet.DeviceIdentifier("1")
	mockClient.On("HeadscaleServiceGetRoutes", mock.Anything).Return(nil, errors.New("test error"))

	_, err := g.getRoutesForDevice(ctx, testID)

	assert.Error(t, err)
}

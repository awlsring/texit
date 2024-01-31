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

func TestCreatePreauthKey(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	testKey := "test-key"
	mockResponse := &headscale_service.HeadscaleServiceCreatePreAuthKeyOK{
		Payload: &models.V1CreatePreAuthKeyResponse{
			PreAuthKey: &models.V1PreAuthKey{
				Key: testKey,
			},
		},
	}

	mockClient.EXPECT().HeadscaleServiceCreatePreAuthKey(mock.Anything).Return(mockResponse, nil)

	resultKey, err := g.CreatePreauthKey(ctx, true)

	assert.NoError(t, err)
	assert.Equal(t, tailnet.PreauthKey(testKey), resultKey)
}

func TestCreatePreauthKey_Error(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockClientService_headscale_service(t)
	g := &HeadscaleGateway{
		client: mockClient,
		user:   "test-user",
	}

	mockClient.EXPECT().HeadscaleServiceCreatePreAuthKey(mock.Anything).Return(nil, errors.New("test error"))

	_, err := g.CreatePreauthKey(ctx, true)

	assert.Error(t, err)
}

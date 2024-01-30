package node

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestGetPlatformGateway(t *testing.T) {
	mockRepo := mocks.NewMockNode_repository(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)
	mockPlatformGws := map[string]gateway.Platform{
		"test": mockPlatformGw,
	}

	s := &Service{
		repo:        mockRepo,
		workSvc:     mockWorkSvc,
		platformGws: mockPlatformGws,
	}

	gw, err := s.getPlatformGateway(context.Background(), provider.Identifier("test"))

	assert.NoError(t, err)
	assert.Equal(t, mockPlatformGw, gw)
}

func TestGetPlatformGatewayError(t *testing.T) {
	mockRepo := mocks.NewMockNode_repository(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	mockPlatformGws := map[string]gateway.Platform{}

	s := &Service{
		repo:        mockRepo,
		workSvc:     mockWorkSvc,
		platformGws: mockPlatformGws,
	}

	gw, err := s.getPlatformGateway(context.Background(), provider.Identifier("test"))

	assert.Error(t, err)
	assert.Nil(t, gw)
	assert.Equal(t, service.ErrUnknownPlatform, errors.Cause(err))
}

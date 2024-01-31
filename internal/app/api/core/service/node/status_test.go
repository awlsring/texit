package node

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStatus(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{"test": mockPlatformGw},
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "test",
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)
	mockPlatformGw.On("GetStatus", mock.Anything, testNode).Return(node.StatusRunning, nil)

	status, err := s.Status(ctx, testNode.Identifier)

	assert.NoError(t, err)
	assert.Equal(t, node.StatusRunning, status)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "GetStatus", ctx, testNode)
}

func TestStatusNodeError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)

	s := &Service{
		repo: mockRepo,
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(nil, errors.New("test error"))

	status, err := s.Status(ctx, testNode.Identifier)

	assert.Error(t, err)
	assert.Equal(t, node.StatusUnknown, status)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestStatusPlatformGatewayError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{},
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "unknown",
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)

	status, err := s.Status(ctx, testNode.Identifier)

	assert.Error(t, err)
	assert.Equal(t, node.StatusUnknown, status)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestStatusPlatformError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{"test": mockPlatformGw},
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "test",
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)
	mockPlatformGw.On("GetStatus", mock.Anything, testNode).Return(node.StatusUnknown, errors.New("test error"))

	status, err := s.Status(ctx, testNode.Identifier)

	assert.Error(t, err)
	assert.Equal(t, node.StatusUnknown, status)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "GetStatus", ctx, testNode)
}

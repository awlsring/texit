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

func TestStart(t *testing.T) {
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
	mockPlatformGw.On("StartNode", mock.Anything, testNode).Return(nil)

	err := s.Start(ctx, testNode.Identifier)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "StartNode", ctx, testNode)
}

func TestStartError(t *testing.T) {
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

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(nil, errors.New("test error"))

	err := s.Start(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestStartPlatformGatewayError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{"test": mockPlatformGw},
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "unknown",
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)

	err := s.Start(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestStartNodeError(t *testing.T) {
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
	mockPlatformGw.On("StartNode", mock.Anything, testNode).Return(errors.New("test error"))

	err := s.Start(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "StartNode", ctx, testNode)
}

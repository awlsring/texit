package node

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStop(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{"test": mockPlatformGw},
		workSvc:     mockWorkSvc,
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "test",
		Ephemeral:  true,
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)
	mockPlatformGw.On("StopNode", mock.Anything, testNode).Return(nil)
	mockWorkSvc.On("LaunchDeprovisionNodeWorkflow", mock.Anything, testNode.Identifier).Return(workflow.ExecutionIdentifier("test"), nil)

	err := s.Stop(ctx, testNode.Identifier)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "StopNode", ctx, testNode)
	mockWorkSvc.AssertCalled(t, "LaunchDeprovisionNodeWorkflow", ctx, testNode.Identifier)
}

func TestStopEphemeral(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{"test": mockPlatformGw},
		workSvc:     mockWorkSvc,
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "test",
		Ephemeral:  false,
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)
	mockPlatformGw.On("StopNode", mock.Anything, testNode).Return(nil)

	err := s.Stop(ctx, testNode.Identifier)

	assert.NoError(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "StopNode", ctx, testNode)
}

func TestStopError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)

	s := &Service{
		repo: mockRepo,
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(nil, errors.New("test error"))

	err := s.Stop(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestStopPlatformGatewayError(t *testing.T) {
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

	err := s.Stop(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestStopNodeError(t *testing.T) {
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
	mockPlatformGw.On("StopNode", mock.Anything, testNode).Return(errors.New("test error"))

	err := s.Stop(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "StopNode", ctx, testNode)
}

func TestStopWorkflowError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)

	s := &Service{
		repo:        mockRepo,
		platformGws: map[string]gateway.Platform{"test": mockPlatformGw},
		workSvc:     mockWorkSvc,
	}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
		Provider:   "test",
		Ephemeral:  true,
	}

	mockRepo.On("Get", mock.Anything, testNode.Identifier).Return(testNode, nil)
	mockPlatformGw.On("StopNode", mock.Anything, testNode).Return(nil)
	mockWorkSvc.On("LaunchDeprovisionNodeWorkflow", mock.Anything, testNode.Identifier).Return(workflow.ExecutionIdentifier(""), errors.New("test error"))

	err := s.Stop(ctx, testNode.Identifier)

	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
	mockPlatformGw.AssertCalled(t, "StopNode", ctx, testNode)
	mockWorkSvc.AssertCalled(t, "LaunchDeprovisionNodeWorkflow", ctx, testNode.Identifier)
}

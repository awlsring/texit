package workflow

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLaunchDeprovisionNodeWorkflow(t *testing.T) {
	ctx := context.Background()
	id := node.Identifier("test-node")

	mockNodeRepo := mocks.NewMockNode_repository(t)
	mockExecRepo := mocks.NewMockExecution_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)
	pGateways := map[string]gateway.Platform{
		"test-provider": mockPlatformGw,
	}
	mockTailnetGw := mocks.NewMockTailnet_gateway(t)
	tGateways := map[string]gateway.Tailnet{
		"test-tailnet": mockTailnetGw,
	}

	s := NewService(mockNodeRepo, mockExecRepo, tGateways, pGateways)

	mockNode := &node.Node{
		Identifier:        id,
		Provider:          "test-provider",
		Tailnet:           "test-tailnet",
		TailnetIdentifier: tailnet.DeviceIdentifier("test-tailnet-identifier"),
	}

	mockNodeRepo.EXPECT().Get(mock.Anything, id).Return(mockNode, nil)
	mockExecRepo.EXPECT().CreateExecution(mock.Anything, mock.Anything).Return(nil)
	mockExecRepo.EXPECT().CloseExecution(mock.Anything, mock.Anything, workflow.StatusComplete, mock.Anything).Return(nil)
	mockPlatformGw.EXPECT().DeleteNode(mock.Anything, mockNode).Return(nil)
	mockTailnetGw.EXPECT().DeleteDevice(mock.Anything, mockNode.TailnetIdentifier).Return(nil)
	mockNodeRepo.EXPECT().Delete(mock.Anything, id).Return(nil)

	exId, err := s.LaunchDeprovisionNodeWorkflow(ctx, id)

	// let goroutine run before finishing tests
	time.Sleep(1 * time.Millisecond)

	assert.NoError(t, err)
	assert.NotNil(t, exId)
}

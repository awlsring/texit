package workflow

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLaunchDeprovisionNodeWorkflow(t *testing.T) {
	ctx := context.Background()
	id := node.Identifier("test-node")

	mockNodeRepo := mocks.NewMockNode_repository(t)
	mockPlatformGw := mocks.NewMockPlatform_gateway(t)
	pGateways := map[string]gateway.Platform{
		"test-provider": mockPlatformGw,
	}
	mockTailnetGw := mocks.NewMockTailnet_gateway(t)
	tGateways := map[string]gateway.Tailnet{
		"test-tailnet": mockTailnetGw,
	}

	s := NewService(mockNodeRepo, tGateways, pGateways)

	mockNode := &node.Node{
		Identifier:        id,
		Provider:          "test-provider",
		Tailnet:           "test-tailnet",
		TailnetIdentifier: tailnet.DeviceIdentifier("test-tailnet-identifier"),
	}

	mockNodeRepo.EXPECT().Get(mock.Anything, id).Return(mockNode, nil)
	mockPlatformGw.EXPECT().DeleteNode(mock.Anything, mockNode).Return(nil)
	mockTailnetGw.EXPECT().DeleteDevice(mock.Anything, mockNode.TailnetIdentifier).Return(nil)
	mockNodeRepo.EXPECT().Delete(mock.Anything, id).Return(nil)

	exId, err := s.LaunchDeprovisionNodeWorkflow(ctx, id)

	// let goroutine run before finishing tests
	time.Sleep(1 * time.Millisecond)

	assert.NoError(t, err)
	assert.NotNil(t, exId)
}

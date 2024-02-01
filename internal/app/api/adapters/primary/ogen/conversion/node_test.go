package conversion

import (
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestTranslateNodeStatus(t *testing.T) {
	tests := []struct {
		name   string
		status node.Status
		want   texit.NodeStatus
	}{
		{
			name:   "Running",
			status: node.StatusRunning,
			want:   texit.NodeStatusRunning,
		},
		{
			name:   "Starting",
			status: node.StatusStarting,
			want:   texit.NodeStatusStarting,
		},
		{
			name:   "Stopping",
			status: node.StatusStopping,
			want:   texit.NodeStatusStopping,
		},
		{
			name:   "Stopped",
			status: node.StatusStopped,
			want:   texit.NodeStatusStopped,
		},
		{
			name:   "Unknown",
			status: node.Status(999), // Unknown status
			want:   texit.NodeStatusUnknown,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, TranslateNodeStatus(tt.status))
		})
	}
}

func TestNodeToSummary(t *testing.T) {
	node := &node.Node{
		Identifier:         node.Identifier("test-id"),
		Provider:           provider.Identifier("test-provider"),
		PlatformIdentifier: node.PlatformIdentifier("test-platform-id"),
		Location:           provider.Location("test-location"),
		Tailnet:            tailnet.Identifier("test-tailnet"),
		TailnetName:        tailnet.DeviceName("test-tailnet-name"),
		TailnetIdentifier:  tailnet.DeviceIdentifier("test-tailnet-id"),
		Ephemeral:          true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	summary := texit.NodeSummary{
		Identifier:              "test-id",
		Provider:                "test-provider",
		ProviderNodeIdentifier:  "test-platform-id",
		Location:                "test-location",
		Tailnet:                 "test-tailnet",
		TailnetDeviceName:       "test-tailnet-name",
		TailnetDeviceIdentifier: "test-tailnet-id",
		Ephemeral:               true,
		Created:                 float64(time.Now().Unix()),
		Updated:                 float64(time.Now().Unix()),
	}

	assert.Equal(t, summary, NodeToSummary(node))
}

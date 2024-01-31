package conversion

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func TranslateNodeStatus(status node.Status) teen.NodeStatus {
	switch status {
	case node.StatusRunning:
		return teen.NodeStatus_NODE_STATUS_RUNNING
	case node.StatusStarting:
		return teen.NodeStatus_NODE_STATUS_STARTING
	case node.StatusStopped:
		return teen.NodeStatus_NODE_STATUS_STOPPED
	case node.StatusStopping:
		return teen.NodeStatus_NODE_STATUS_STOPPING
	default:
		return teen.NodeStatus_NODE_STATUS_UNKNOWN_UNSPECIFIED
	}
}

func NodeToSummary(node *node.Node) *teen.NodeSummary {
	return &teen.NodeSummary{
		Id:          node.Identifier.String(),
		Provider:    node.Provider.String(),
		PlatformId:  node.PlatformIdentifier.String(),
		TailnetId:   node.TailnetIdentifier.String(),
		Tailnet:     node.Tailnet.String(),
		TailnetName: node.TailnetName.String(),
		Location:    node.Location.String(),
		Ephemeral:   node.Ephemeral,
		CreatedAt:   node.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:   node.UpdatedAt.Format(time.RFC3339Nano),
	}
}

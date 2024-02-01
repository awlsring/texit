package conversion

import (
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func TranslateNodeStatus(s node.Status) texit.NodeStatus {
	switch s {
	case node.StatusRunning:
		return texit.NodeStatusRunning
	case node.StatusStarting:
		return texit.NodeStatusStarting
	case node.StatusStopping:
		return texit.NodeStatusStopping
	case node.StatusStopped:
		return texit.NodeStatusStopped
	default:
		return texit.NodeStatusUnknown
	}
}

func NodeToSummary(n *node.Node) texit.NodeSummary {
	return texit.NodeSummary{
		Identifier:              n.Identifier.String(),
		Provider:                n.Provider.String(),
		ProviderNodeIdentifier:  n.PlatformIdentifier.String(),
		Location:                n.Location.String(),
		Tailnet:                 n.Tailnet.String(),
		TailnetDeviceName:       n.TailnetName.String(),
		TailnetDeviceIdentifier: n.TailnetIdentifier.String(),
		Ephemeral:               n.Ephemeral,
		Created:                 float64(n.CreatedAt.Unix()),
		Updated:                 float64(n.UpdatedAt.Unix()),
	}
}

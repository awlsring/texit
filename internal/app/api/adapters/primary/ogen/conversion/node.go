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
	case node.StatusPending:
		return texit.NodeStatusPending
	default:
		return texit.NodeStatusUnknown
	}
}

func TranslateNodeProvisioningStatus(s node.ProvisionStatus) texit.ProvisioningStatus {
	switch s {
	case node.ProvisionStatusCreated:
		return texit.ProvisioningStatusCreated
	case node.ProvisionStatusCreating:
		return texit.ProvisioningStatusCreating
	case node.ProvisionStatusFailed:
		return texit.ProvisioningStatusFailed
	default:
		return texit.ProvisioningStatusUnknown
	}
}

func TranslateNodeSize(s node.Size) texit.NodeSize {
	switch s {
	case node.SizeSmall:
		return texit.NodeSizeSmall
	case node.SizeMedium:
		return texit.NodeSizeMedium
	case node.SizeLarge:
		return texit.NodeSizeLarge
	default:
		return texit.NodeSizeUnknown
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
		Size:                    TranslateNodeSize(n.Size),
		Created:                 float64(n.CreatedAt.Unix()),
		Updated:                 float64(n.UpdatedAt.Unix()),
		ProvisioningStatus:      TranslateNodeProvisioningStatus(n.ProvisionStatus),
	}
}

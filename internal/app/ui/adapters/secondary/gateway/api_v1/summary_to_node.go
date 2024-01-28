package apiv1

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func TranslateNodeStatus(s v1.NodeStatus) node.Status {
	switch s {
	case v1.NodeStatus_NODE_STATUS_RUNNING:
		return node.StatusActive
	case v1.NodeStatus_NODE_STATUS_STOPPED, v1.NodeStatus_NODE_STATUS_STOPPING:
		return node.StatusInactive
	default:
		return node.StatusUnknown
	}
}

func SummaryToNode(s *v1.NodeSummary) (*node.Node, error) {
	id, err := node.IdentifierFromString(s.Id)
	if err != nil {
		return nil, err
	}

	provId, err := provider.IdentifierFromString(s.ProviderId)
	if err != nil {
		return nil, err
	}

	platId, err := node.PlatformIdentifierFromString(s.PlatformId)
	if err != nil {
		return nil, err
	}

	tailId, err := tailnet.DeviceIdentifierFromString(s.TailnetId)
	if err != nil {
		return nil, err
	}

	created, err := time.Parse(time.RFC3339, s.CreatedAt)
	if err != nil {
		return nil, err
	}

	updated, err := time.Parse(time.RFC3339, s.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &node.Node{
		Identifier:         id,
		ProviderIdentifier: provId,
		PlatformIdentifier: platId,
		TailnetIdentifier:  tailId,
		Location:           provider.Location(s.Location),
		CreatedAt:          created,
		UpdatedAt:          updated,
	}, nil
}

package conversion

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func NodeToSummary(node *node.Node) *teen.NodeSummary {
	return &teen.NodeSummary{
		Id:          node.Identifier.String(),
		Provider:    node.Provider.String(),
		PlatformId:  node.PlatformIdentifier.String(),
		TailnetId:   node.TailnetName.String(),
		Tailnet:     node.Tailnet.String(),
		TailnetName: node.TailnetName.String(),
		Location:    node.Location.String(),
		Ephemeral:   node.Ephemeral,
		CreatedAt:   node.CreatedAt.Format(time.RFC3339Nano),
		UpdatedAt:   node.UpdatedAt.Format(time.RFC3339Nano),
	}
}

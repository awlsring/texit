package conversion

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func NodeToSummary(node *node.Node) *teen.NodeSummary {
	return &teen.NodeSummary{
		Id:         node.Identifier.String(),
		ProviderId: node.ProviderIdentifier.String(),
		PlatformId: node.PlatformIdentifier.String(),
		TailnetId:  node.TailnetIdentifier.String(),
		Location:   node.Location.String(),
		CreatedAt:  node.CreatedAt.String(),
		UpdatedAt:  node.UpdatedAt.String(),
	}
}

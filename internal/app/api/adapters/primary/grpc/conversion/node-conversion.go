package conversion

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func NodeToSummary(node *node.Node) *teen.NodeSummary {
	return &teen.NodeSummary{} // TODO: Implement
}

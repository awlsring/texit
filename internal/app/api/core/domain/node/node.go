package node

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/tailnet"
)

type Node struct {
	// The internal identifier of the node.
	Identifier Identifier
	// The platform identifier of the node.
	PlatformIdentifier PlatformIdentifier
	// The provider the node belong to.
	ProviderIdentifier provider.Identifier
	// The identifier of the node on the tailnet
	TailnetIdentifier tailnet.DeviceIdentifier
	// The location the node is in
	Location provider.Location
	// The preauthkey used to create the node
	PreauthKey tailnet.PreauthKey
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

package node

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

type Node struct {
	// The internal identifier of the node.
	Identifier Identifier
	// The platform identifier of the node.
	PlatformIdentifier PlatformIdentifier
	// The provider the node belong to.
	Provider provider.Identifier
	// The tailnet the node belongs to.
	Tailnet tailnet.Identifier
	// The name of the node on the tailnet
	TailnetName tailnet.DeviceName
	// The id of the node on the tailnet
	TailnetIdentifier tailnet.DeviceIdentifier
	// The location the node is in
	Location provider.Location
	// The preauthkey used to create the node
	PreauthKey tailnet.PreauthKey
	// if the node will be deleted when stopped
	Ephemeral bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

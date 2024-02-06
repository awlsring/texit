package node

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

type Node struct {
	Identifier         node.Identifier
	Status             node.Status
	PlatformIdentifier node.PlatformIdentifier
	Provider           provider.Identifier
	ProviderType       provider.Type
	Tailnet            tailnet.Identifier
	TailnetType        tailnet.Type
	TailnetName        tailnet.DeviceName
	TailnetIdentifier  tailnet.DeviceIdentifier
	Location           provider.Location
	Ephemeral          bool
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

func NewBaseNode(n *node.Node) *Node {
	return &Node{
		Identifier:         n.Identifier,
		Status:             node.StatusUnknown,
		PlatformIdentifier: n.PlatformIdentifier,
		Provider:           n.Provider,
		ProviderType:       provider.TypeUnknown,
		Tailnet:            n.Tailnet,
		TailnetType:        tailnet.TypeUnknown,
		TailnetName:        n.TailnetName,
		TailnetIdentifier:  n.TailnetIdentifier,
		Location:           n.Location,
		Ephemeral:          n.Ephemeral,
		CreatedAt:          n.CreatedAt,
		UpdatedAt:          n.UpdatedAt,
	}
}

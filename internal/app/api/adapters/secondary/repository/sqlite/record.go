package sqlite_node_repository

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
)

type NodeSqlRecord struct {
	Identifier         string    `db:"identifier"`
	PlatformIdentifier string    `db:"platform_identifier"`
	ProviderIdentifier string    `db:"provider_identifier"`
	TailnetIdentifier  string    `db:"tailnet_identifier"`
	Location           string    `db:"location"`
	PreauthKey         string    `db:"preauth_key"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func (n *NodeSqlRecord) ToNode() *node.Node {
	return &node.Node{
		Identifier:         node.Identifier(n.Identifier),
		PlatformIdentifier: node.PlatformIdentifier(n.PlatformIdentifier),
		ProviderIdentifier: provider.Identifier(n.ProviderIdentifier),
		TailnetIdentifier:  tailnet.DeviceIdentifier(n.TailnetIdentifier),
		Location:           provider.Location(n.Location),
		PreauthKey:         tailnet.PreauthKey(n.PreauthKey),
		CreatedAt:          n.CreatedAt,
		UpdatedAt:          n.UpdatedAt,
	}
}

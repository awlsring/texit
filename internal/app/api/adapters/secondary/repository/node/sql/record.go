package sql_node_repository

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

type NodeSqlRecord struct {
	Identifier         string    `db:"identifier"`
	PlatformIdentifier string    `db:"platform_identifier"`
	Provider           string    `db:"provider_identifier"`
	Tailnet            string    `db:"tailnet"`
	TailnetName        string    `db:"tailnet_device_name"`
	TailnetIdentifier  string    `db:"tailnet_identifier"`
	Location           string    `db:"location"`
	PreauthKey         string    `db:"preauth_key"`
	Ephemeral          bool      `db:"ephemeral"`
	Size               string    `db:"size"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

func (n *NodeSqlRecord) ToNode() *node.Node {
	s, _ := node.SizeFromString(n.Size)
	return &node.Node{
		Identifier:         node.Identifier(n.Identifier),
		PlatformIdentifier: node.PlatformIdentifier(n.PlatformIdentifier),
		Provider:           provider.Identifier(n.Provider),
		Tailnet:            tailnet.Identifier(n.Tailnet),
		TailnetIdentifier:  tailnet.DeviceIdentifier(n.TailnetIdentifier),
		TailnetName:        tailnet.DeviceName(n.TailnetName),
		Location:           provider.Location(n.Location),
		PreauthKey:         tailnet.PreauthKey(n.PreauthKey),
		Ephemeral:          n.Ephemeral,
		Size:               s,
		CreatedAt:          n.CreatedAt,
		UpdatedAt:          n.UpdatedAt,
	}
}

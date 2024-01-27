package service

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
)

type Workflow interface {
	ProvisionNode(context.Context, provider.Identifier, provider.Location) (*node.Node, error)
	DeprovisionNode(context.Context, node.Identifier) error
}

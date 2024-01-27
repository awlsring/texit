package service

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
)

type Node interface {
	Create(ctx context.Context) (*node.Node, error)
	Delete(ctx context.Context, id node.Identifier) error
	Start(ctx context.Context, id node.Identifier) error
	Stop(ctx context.Context, id node.Identifier) error
	Describe(ctx context.Context, id node.Identifier) (*node.Node, error)
	List(ctx context.Context) ([]*node.Node, error)
}

package repository

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
)

var (
	ErrNodeNotFound = errors.New("node not found")
)

type Node interface {
	Init(ctx context.Context) error
	Get(ctx context.Context, id node.Identifier) (*node.Node, error)
	List(ctx context.Context) ([]*node.Node, error)
	Delete(ctx context.Context, id node.Identifier) error
	Create(ctx context.Context, node *node.Node) error
}

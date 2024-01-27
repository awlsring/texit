package service

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
)

var (
	ErrUnknownPlatform = errors.New("unknown platform")
)

type Node interface {
	Create(context.Context, provider.Identifier, provider.Location) (*node.Node, error)
	Delete(ctx context.Context, id node.Identifier) error
	Start(ctx context.Context, id node.Identifier) error
	Stop(ctx context.Context, id node.Identifier) error
	Status(ctx context.Context, id node.Identifier) (node.Status, error)
	Describe(ctx context.Context, id node.Identifier) (*node.Node, error)
	List(ctx context.Context) ([]*node.Node, error)
}

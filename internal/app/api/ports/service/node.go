package service

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/node"
)

var (
	ErrUnknownTailnetId = errors.New("unknown tailnet")
	ErrUnknownPlatform  = errors.New("unknown platform")
	ErrUnknownNode      = errors.New("unknown node")
)

type Node interface {
	Start(ctx context.Context, id node.Identifier) error
	Stop(ctx context.Context, id node.Identifier) error
	Status(ctx context.Context, id node.Identifier) (node.Status, error)
	Describe(ctx context.Context, id node.Identifier) (*node.Node, error)
	List(ctx context.Context) ([]*node.Node, error)
}

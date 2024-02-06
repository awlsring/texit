package service

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/core/domain/node"
	cnode "github.com/awlsring/texit/internal/pkg/domain/node"
)

type Node interface {
	DescribeNode(context.Context, cnode.Identifier) (*node.Node, error)
	ListNodes(context.Context) ([]*node.Node, error)
}

package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (s *Service) GetNode(ctx context.Context, node node.Identifier) (*node.Node, error) {
	return s.apiGw.GetNode(ctx, node)
}

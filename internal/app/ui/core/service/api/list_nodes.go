package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (s *Service) ListNodes(ctx context.Context) ([]*node.Node, error) {
	return s.apiGw.ListNodes(ctx)
}

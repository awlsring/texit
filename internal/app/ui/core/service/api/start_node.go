package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (s *Service) StartNode(ctx context.Context, node node.Identifier) error {
	return s.apiGw.StartNode(ctx, node)
}

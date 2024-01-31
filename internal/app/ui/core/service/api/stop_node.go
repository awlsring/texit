package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (s *Service) StopNode(ctx context.Context, node node.Identifier) error {
	return s.apiGw.StopNode(ctx, node)
}

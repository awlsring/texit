package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
)

func (s *Service) GetNodeStatus(ctx context.Context, node node.Identifier) (node.
	Status, error) {
	return s.apiGw.GetNodeStatus(ctx, node)
}

package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

func (s *Service) DeprovisionNode(ctx context.Context, node node.Identifier) (workflow.ExecutionIdentifier, error) {
	return s.apiGw.DeprovisionNode(ctx, node)
}

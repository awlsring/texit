package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/pkg/errors"
)

func (s *Service) DeprovisionNode(ctx context.Context, node node.Identifier) (workflow.ExecutionIdentifier, error) {
	resp, err := s.apiGw.DeprovisionNode(ctx, node)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return "", service.ErrUnknownNode
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return "", service.ErrInvalidInputError
		}
		return "", err
	}
	return resp, nil
}

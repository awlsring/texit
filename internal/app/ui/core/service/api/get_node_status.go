package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/pkg/errors"
)

func (s *Service) GetNodeStatus(ctx context.Context, n node.Identifier) (node.
	Status, error) {
	resp, err := s.apiGw.GetNodeStatus(ctx, n)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return node.StatusUnknown, service.ErrUnknownNode
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return node.StatusUnknown, service.ErrInvalidInputError
		}
		return node.StatusUnknown, err
	}
	return resp, nil
}

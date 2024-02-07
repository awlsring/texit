package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/pkg/errors"
)

func (s *Service) StartNode(ctx context.Context, n node.Identifier) error {
	err := s.apiGw.StartNode(ctx, n)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return service.ErrUnknownNode
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return service.ErrInvalidInputError
		}
		return err
	}
	return nil
}

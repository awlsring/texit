package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/pkg/errors"
)

func (s *Service) GetNode(ctx context.Context, n node.Identifier) (*node.Node, error) {
	resp, err := s.apiGw.DescribeNode(ctx, n)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return nil, service.ErrUnknownNode
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return nil, service.ErrInvalidInputError

		}
		return nil, err
	}
	return resp, nil
}

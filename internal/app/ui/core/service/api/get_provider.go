package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/pkg/errors"
)

func (s *Service) GetProvider(ctx context.Context, identifier provider.Identifier) (*provider.Provider, error) {
	resp, err := s.apiGw.DescribeProvider(ctx, identifier)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return nil, service.ErrUnknownProvider
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return nil, service.ErrInvalidInputError

		}
		return nil, err
	}
	return resp, nil
}

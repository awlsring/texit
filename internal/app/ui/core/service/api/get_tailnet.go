package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/pkg/errors"
)

func (s *Service) GetTailnet(ctx context.Context, identifier tailnet.Identifier) (*tailnet.Tailnet, error) {
	resp, err := s.apiGw.DescribeTailnet(ctx, identifier)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return nil, service.ErrUnknownTailnet
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return nil, service.ErrInvalidInputError

		}
		return nil, err
	}
	return resp, nil
}

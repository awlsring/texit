package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

func (s *Service) GetTailnet(ctx context.Context, identifier tailnet.Identifier) (*tailnet.Tailnet, error) {
	return s.apiGw.DescribeTailnet(ctx, identifier)
}

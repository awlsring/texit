package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

func (s *Service) ListTailnets(ctx context.Context) ([]*tailnet.Tailnet, error) {
	return s.apiGw.ListTailnets(ctx)
}

package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

func (s *Service) ListProviders(ctx context.Context) ([]*provider.Provider, error) {
	return s.apiGw.ListProviders(ctx)
}

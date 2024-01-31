package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

func (s *Service) GetProvider(ctx context.Context, identifier provider.Identifier) (*provider.Provider, error) {
	return s.apiGw.GetProvider(ctx, identifier)
}

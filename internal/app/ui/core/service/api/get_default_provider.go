package api

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
)

func (s *Service) GetDefaultProvider(ctx context.Context) (*provider.Provider, error) {
	provider, err := s.apiGw.GetDefaultProvider(ctx)
	if err != nil {
		return nil, err
	}

	return provider, nil
}

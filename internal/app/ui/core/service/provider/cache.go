package provider

import (
	"context"
	"errors"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

const (
	defaultExpiration = 5 * time.Minute
)

func (s *Service) refreshProviders(ctx context.Context) error {
	s.lastRefresh = time.Now()
	providers, err := s.apiGw.ListProviders(ctx)
	if err != nil {
		return err
	}
	for _, p := range providers {
		s.providers[string(p.Name.String())] = p
	}
	return nil
}

func (s *Service) getProviderFromMap(ctx context.Context, id provider.Identifier) (*provider.Provider, error) {
	s.mut.Lock()
	if time.Since(s.lastRefresh) > defaultExpiration {
		if err := s.refreshProviders(ctx); err != nil {
			return nil, err
		}
	}
	s.mut.Unlock()
	p, ok := s.providers[string(id.String())]
	if !ok {
		return nil, errors.New("provider not found")
	}
	return p, nil
}

package service

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
)

type Api interface {
	GetDefaultProvider(ctx context.Context) (*provider.Provider, error)
	GetProvider(ctx context.Context, id provider.Identifier) (*provider.Provider, error)
	ListProviders(ctx context.Context) ([]*provider.Provider, error)
	CheckServerHealth(ctx context.Context) error
}

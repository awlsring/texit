package service

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
)

type Provider interface {
	Describe(context.Context, provider.Identifier) (*provider.Provider, error)
	List(context.Context) ([]*provider.Provider, error)
}

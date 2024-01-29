package service

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
)

var (
	ErrUnknownProvider = errors.New("unknown provider")
	ErrNoProviders     = errors.New("no providers")
)

type Provider interface {
	Describe(context.Context, provider.Identifier) (*provider.Provider, error)
	List(context.Context) ([]*provider.Provider, error)
}

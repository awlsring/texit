package service

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
)

var (
	ErrUnknownProvider         = errors.New("unknown provider")
	ErrMulitpleDefaultProvider = errors.New("multiple default providers")
	ErrNoProviders             = errors.New("no providers")
)

type Provider interface {
	Default(context.Context) (*provider.Provider, error)
	Describe(context.Context, provider.Identifier) (*provider.Provider, error)
	List(context.Context) ([]*provider.Provider, error)
}

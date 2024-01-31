package service

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

var (
	ErrUnknownProvider = errors.New("unknown provider")
	ErrNoProviders     = errors.New("no providers")
)

type Provider interface {
	Describe(context.Context, provider.Identifier) (*provider.Provider, error)
	List(context.Context) ([]*provider.Provider, error)
}

package service

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

var (
	ErrUnknownProvider = errors.New("unknown provider")
)

type Provider interface {
	DescribeProvider(context.Context, provider.Identifier) (*provider.Provider, error)
	ListProviders(context.Context) ([]*provider.Provider, error)
}

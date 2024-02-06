package service

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

type Provider interface {
	DescribeProvider(context.Context, provider.Identifier) (*provider.Provider, error)
	ListProviders(context.Context) ([]*provider.Provider, error)
}

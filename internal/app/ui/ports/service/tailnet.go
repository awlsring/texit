package service

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

type Tailnet interface {
	DescribeTailnet(context.Context, tailnet.Identifier) (*tailnet.Tailnet, error)
	ListTailnets(context.Context) ([]*tailnet.Tailnet, error)
}

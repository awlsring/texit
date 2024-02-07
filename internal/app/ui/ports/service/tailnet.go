package service

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

var (
	ErrUnknownTailnet = errors.New("unknown tailnet")
)

type Tailnet interface {
	DescribeTailnet(context.Context, tailnet.Identifier) (*tailnet.Tailnet, error)
	ListTailnets(context.Context) ([]*tailnet.Tailnet, error)
}

package service

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

var (
	ErrUnknownTailnet = errors.New("unknown provider")
	ErrNoTailnets     = errors.New("no providers")
)

type Tailnet interface {
	Describe(context.Context, tailnet.Identifier) (*tailnet.Tailnet, error)
	List(context.Context) ([]*tailnet.Tailnet, error)
}

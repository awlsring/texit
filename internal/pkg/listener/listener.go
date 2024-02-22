package listener

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/rs/zerolog"
)

type ListenerOption func(Listener)

func WithLogger(logger zerolog.Logger) ListenerOption {
	return func(l Listener) {
		l.SetLogger(logger)
	}
}

func WithLogLevel(level zerolog.Level) ListenerOption {
	return func(l Listener) {
		l.SetLogLevel(level)
	}
}

type Handler interface {
	Handle(context.Context, notification.ExecutionMessage)
}

type Listener interface {
	Subscribe(context.Context, string) error
	SetLogLevel(zerolog.Level)
	SetLogger(zerolog.Logger)
}

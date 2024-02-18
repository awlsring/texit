package mqtt

import "github.com/rs/zerolog"

type ListenerOption func(*listener)

func WithLogger(logger zerolog.Logger) ListenerOption {
	return func(l *listener) {
		l.log = logger
	}
}

func WithLogLevel(level zerolog.Level) ListenerOption {
	return func(l *listener) {
		l.lvl = level
	}
}

package logger

import (
	"context"
	"os"
	"runtime/debug"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var DefaultLogger zerolog.Logger = log.With().Logger()

type LoggingOpt func(zctx zerolog.Context, ctx context.Context) zerolog.Context

func WithField(key string, value interface{}) LoggingOpt {
	return func(zctx zerolog.Context, ctx context.Context) zerolog.Context {
		return zctx.Interface(key, value)
	}
}

func InitContextLogger(ctx context.Context, lvl zerolog.Level, opts ...LoggingOpt) context.Context {
	buildInfo, _ := debug.ReadBuildInfo()
	logger := zerolog.New(os.Stderr)
	logger = logger.Level(lvl)
	logger = logger.With().
		Caller().
		Timestamp().
		Caller().
		Int("pid", os.Getpid()).
		Str("go_version", buildInfo.GoVersion).
		Logger()

	for _, opt := range opts {
		logger.UpdateContext(func(zctx zerolog.Context) zerolog.Context {
			return opt(zctx, ctx)
		})
	}

	ctx = logger.WithContext(ctx)

	return ctx
}

func FromContext(ctx context.Context) zerolog.Logger {
	logger := log.Ctx(ctx)
	if logger.GetLevel() == zerolog.Disabled {
		return DefaultLogger
	}
	return *logger
}

func UsePrettyLogger(log zerolog.Logger) zerolog.Logger {
	return log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

func LogLevelFromEnv() zerolog.Level {
	lvlString := os.Getenv("LOG_LEVEL")
	if lvlString == "" {
		return zerolog.InfoLevel
	}
	lvl, err := zerolog.ParseLevel(lvlString)
	if err != nil {
		return zerolog.InfoLevel
	}
	return lvl
}

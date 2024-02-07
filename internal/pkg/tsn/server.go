package tsn

import (
	"github.com/rs/zerolog"
	"tailscale.com/tsnet"
)

type ServerOption func(*tsnet.Server)

func WithStandardLoggingFunc(log zerolog.Logger) ServerOption {
	tailog := log.With().Timestamp().Str("process", "tsnet").Logger()
	return func(s *tsnet.Server) {
		s.Logf = func(format string, args ...interface{}) {
			tailog.Debug().Msgf(format, args...)
		}
	}
}

func WithAuthKey(authKey string) ServerOption {
	return func(s *tsnet.Server) {
		s.AuthKey = authKey
	}
}

func WithHostname(hostname string) ServerOption {
	return func(s *tsnet.Server) {
		s.Hostname = hostname
	}
}

func WithStateDir(stateDir string) ServerOption {
	return func(s *tsnet.Server) {
		s.Dir = stateDir
	}
}

func WithControlURL(controlURL string) ServerOption {
	return func(s *tsnet.Server) {
		s.ControlURL = controlURL
	}
}

func WithRunWebClient(runWebClient bool) ServerOption {
	return func(s *tsnet.Server) {
		s.RunWebClient = runWebClient
	}
}

func WithLogFunc(logFunc func(format string, args ...interface{})) ServerOption {
	return func(s *tsnet.Server) {
		s.Logf = logFunc
	}
}

func NewServer(options ...ServerOption) *tsnet.Server {
	s := &tsnet.Server{
		Logf: func(format string, args ...interface{}) {},
	}

	for _, opt := range options {
		opt(s)
	}
	return s
}

package ogen

import (
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/rs/zerolog"
)

type ServerOpt func(*Server)

func WithSecurityHandler(h texit.SecurityHandler) ServerOpt {
	return func(s *Server) {
		s.secHdl = h
	}
}

func WithServerOptions(opts ...texit.ServerOption) ServerOpt {
	return func(s *Server) {
		s.opts = append(s.opts, opts...)
	}
}

func WithLogLevel(level zerolog.Level) ServerOpt {
	return func(s *Server) {
		s.logLevel = level
	}
}

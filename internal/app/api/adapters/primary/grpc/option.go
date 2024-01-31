package grpc

import (
	"github.com/rs/zerolog"
)

type ServerOpt func(*GrpcServer)

func WithStaticKey(key string) ServerOpt {
	return func(s *GrpcServer) {
		s.apiKey = key
	}
}

func WithLogLevel(level zerolog.Level) ServerOpt {
	return func(s *GrpcServer) {
		s.logLevel = level
	}
}

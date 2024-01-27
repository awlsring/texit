package grpc

import "github.com/rs/zerolog"

type ServerOpt func(*GrpcServer)

func WithNetwork(network string) ServerOpt {
	return func(s *GrpcServer) {
		s.network = network
	}
}

func WithAddress(address string) ServerOpt {
	return func(s *GrpcServer) {
		s.address = address
	}
}

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

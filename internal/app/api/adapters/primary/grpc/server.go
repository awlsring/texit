package grpc

import (
	"context"
	"net"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/interceptor"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

const (
	defaultStaticKey = "set-me"
	defaultLogLevel  = zerolog.InfoLevel
)

type GrpcServer struct {
	apiKey   string
	logLevel zerolog.Level
	srv      *grpc.Server
	listener net.Listener
}

func NewServer(lis net.Listener, hdl teen.TailscaleEphemeralExitNodesServiceServer, opts ...ServerOpt) (*GrpcServer, error) {
	s := &GrpcServer{
		apiKey:   defaultStaticKey,
		logLevel: defaultLogLevel,
		listener: lis,
	}

	for _, opt := range opts {
		opt(s)
	}

	grpcOpts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(interceptor.NewLoggingInterceptor(s.logLevel), interceptor.NewStaticAuthKeyInterceptor(s.apiKey)),
	}
	s.srv = grpc.NewServer(grpcOpts...)
	teen.RegisterTailscaleEphemeralExitNodesServiceServer(s.srv, hdl)
	return s, nil
}

func (s *GrpcServer) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting server...")

	go func() {
		log.Debug().Msgf("server listening at %v", s.listener.Addr())
		if err := s.srv.Serve(s.listener); err != nil {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	go func() {
		<-ctx.Done()
		log.Debug().Msg("Shutting down server...")
		s.srv.GracefulStop()
	}()

	return nil
}

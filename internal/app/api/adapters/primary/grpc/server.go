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
	defaultNetwork  = "tcp"
	defaultAddress  = ":7032"
	defaultLogLevel = zerolog.InfoLevel
)

type GrpcServer struct {
	network  string
	address  string
	logLevel zerolog.Level
	srv      *grpc.Server
	listener net.Listener
}

func NewServer(hdl teen.TailscaleEphemeralExitNodesServiceServer, opts ...ServerOpt) (*GrpcServer, error) {
	s := &GrpcServer{
		network:  defaultNetwork,
		address:  defaultAddress,
		logLevel: defaultLogLevel,
	}

	for _, opt := range opts {
		opt(s)
	}

	grpcOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.NewLoggingInterceptor(s.logLevel)),
	}
	s.srv = grpc.NewServer(grpcOpts...)
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return nil, err
	}
	s.listener = lis

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

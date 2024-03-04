package ogen

import (
	"context"
	"net"
	"net/http"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/middleware"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/smithy_errors"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

const (
	defaultLogLevel = zerolog.InfoLevel
)

type Server struct {
	logLevel zerolog.Level
	hdl      texit.Handler
	secHdl   texit.SecurityHandler
	opts     []texit.ServerOption
	srv      http.Handler
}

func NewServer(hdl texit.Handler, opts ...ServerOpt) (*Server, error) {
	s := &Server{
		logLevel: defaultLogLevel,
		hdl:      hdl,
		opts: []texit.ServerOption{
			texit.WithMiddleware(middleware.LoggingMiddleware(zerolog.DebugLevel)),
			texit.WithNotFound(smithy_errors.UnknownOperationHandler),
			texit.WithErrorHandler(smithy_errors.ResponseHandlerWithLogger(zerolog.DebugLevel)),
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	srv, err := texit.NewServer(s.hdl, s.secHdl, s.opts...)
	if err != nil {
		return nil, err
	}
	s.srv = srv

	return s, nil
}

func (s *Server) HttpHandler() http.Handler {
	return s.srv
}

func (s *Server) Serve(ctx context.Context, lis net.Listener) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting server...")

	go func() {
		log.Debug().Msgf("server listening at %v", lis.Addr())
		if err := http.Serve(lis, s.srv); err != nil {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	go func() {
		log.Debug().Msgf("metrics listening at %v", ":9090")
		if err := http.ListenAndServe(":9090", promhttp.Handler()); err != nil {
			log.Error().Err(err).Msg("Metrics error")
		}
	}()

	go func() {
		<-ctx.Done()
		log.Debug().Msg("Shutting down server...")
	}()

	return nil
}

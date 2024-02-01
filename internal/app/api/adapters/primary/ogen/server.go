package ogen

import (
	"context"
	"net"
	"net/http"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/middleware"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/smithy_errors"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/rs/zerolog"
)

const (
	defaultLogLevel = zerolog.InfoLevel
)

type Server struct {
	logLevel zerolog.Level
	hdl      texit.Handler
	lis      net.Listener
	secHdl   texit.SecurityHandler
	opts     []texit.ServerOption
}

func NewServer(lis net.Listener, hdl texit.Handler, opts ...ServerOpt) *Server {
	s := &Server{
		logLevel: defaultLogLevel,
		hdl:      hdl,
		lis:      lis,
		opts: []texit.ServerOption{
			texit.WithMiddleware(middleware.LoggingMiddleware(zerolog.DebugLevel)),
			texit.WithNotFound(smithy_errors.UnknownOperationHandler),
			texit.WithErrorHandler(smithy_errors.ResponseHandlerWithLogger(zerolog.DebugLevel)),
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Starting server...")

	srv, err := texit.NewServer(s.hdl, s.secHdl, s.opts...)
	if err != nil {
		log.Error().Err(err).Msg("Error creating server")
		return err
	}

	go func() {
		log.Debug().Msgf("server listening at %v", s.lis.Addr())
		if err := http.Serve(s.lis, srv); err != nil {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	go func() {
		<-ctx.Done()
		log.Debug().Msg("Shutting down server...")
		s.lis.Close()
	}()

	return nil
}

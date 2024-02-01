package smithy_errors

import (
	"context"
	"net/http"

	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

const (
	SmithyErrorTypeHeader = "X-Amzn-Errortype"
)

func ResponseHandlerWithLogger(lvl zerolog.Level) func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
		ctx = logger.InitContextLogger(ctx, lvl)
		ResponseHandler(ctx, w, r, err)
	}
}

func ResponseHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) { //
	log := logger.FromContext(ctx)
	log.Debug().Err(err).Msg("Handling error as smithy exception")

	log.Debug().Msg("Getting exception from error")
	serr := translateError(err)
	log.Debug().Interface("exception", err).Msgf("got exception")

	log.Debug().Msg("Building response")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(SmithyErrorTypeHeader, serr.Type().String())
	w.WriteHeader(serr.Code())
	_, e := w.Write(serr.AsJsonMessage())
	log.Error().Err(e).Msg("Failed to write response")
}

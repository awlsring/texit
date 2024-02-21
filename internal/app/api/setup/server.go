package setup

import (
	"net"

	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/tsn"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LoadListener(cfg *config.ServerConfig) net.Listener {
	if cfg.Tailnet != nil {
		l, err := tsn.ListenerFromConfig(*cfg.Tailnet, cfg.Address, tsn.WithStandardLoggingFunc(zerolog.DebugLevel))
		appinit.PanicOnErr(err)
		return l
	}
	log.Info().Msg("Creating normal net listener")
	l, err := net.Listen("tcp", cfg.Address)
	appinit.PanicOnErr(err)
	return l
}

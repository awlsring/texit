package discord

import (
	"context"
	"net"

	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/tsn"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/rs/zerolog/log"
)

type Sec struct {
	key string
}

func (s Sec) SmithyAPIHttpApiKeyAuth(ctx context.Context, operationName string) (texit.SmithyAPIHttpApiKeyAuth, error) {
	return texit.SmithyAPIHttpApiKeyAuth{
		APIKey: s.key,
	}, nil
}

func LoadTexitClient(address string, key string) texit.Invoker {
	c, err := texit.NewClient(address, Sec{key: key})
	appinit.PanicOnErr(err)
	return c
}

func LoadListener(cfg discfg.ServerConfig) net.Listener {
	if cfg.Tailnet != nil {
		l, err := tsn.ListenerFromConfig(*cfg.Tailnet, cfg.Address)
		appinit.PanicOnErr(err)
		return l
	}
	log.Info().Msg("Creating net listener")
	l, err := net.Listen("tcp", cfg.Address)
	appinit.PanicOnErr(err)
	return l
}

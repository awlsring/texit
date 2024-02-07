package main

import (
	"context"
	"net"
	"os"
	"os/signal"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord"
	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	api_gateway "github.com/awlsring/texit/internal/app/ui/adapters/secondary/gateway/api"
	"github.com/awlsring/texit/internal/app/ui/config"
	"github.com/awlsring/texit/internal/app/ui/core/service/api"
	"github.com/awlsring/texit/internal/app/ui/core/service/node"
	"github.com/awlsring/texit/internal/app/ui/core/service/provider"
	"github.com/awlsring/texit/internal/app/ui/core/service/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/tsn"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/rs/zerolog"
)

const (
	configEnvVar          = "DISCORD_CONFIG_PATH"
	defaultConfigLocation = "/etc/texit_discord/config.yaml"
)

var log zerolog.Logger

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getConfigPath() string {
	path := os.Getenv(configEnvVar)
	if path == "" {
		return defaultConfigLocation
	}
	return path
}

type Sec struct {
	key string
}

func (s Sec) SmithyAPIHttpApiKeyAuth(ctx context.Context, operationName string) (texit.SmithyAPIHttpApiKeyAuth, error) {
	return texit.SmithyAPIHttpApiKeyAuth{
		APIKey: s.key,
	}, nil
}

func initClient(address string, key string) texit.Invoker {
	c, err := texit.NewClient(address, Sec{key: key})
	panicOnErr(err)
	return c
}

func initListener(cfg discfg.ServerConfig) net.Listener {
	if cfg.Tailnet != nil {
		l, err := tsn.ListenerFromConfig(*cfg.Tailnet, cfg.Address, tsn.WithStandardLoggingFunc(log))
		panicOnErr(err)
		return l
	}
	log.Info().Msg("Creating net listener")
	l, err := net.Listen("tcp", cfg.Address)
	panicOnErr(err)
	return l
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = logger.InitContextLogger(ctx, zerolog.DebugLevel)
	log = logger.FromContext(ctx)
	log.Info().Msg("Initializing")

	log.Info().Msg("Loading config")
	cfg, err := config.LoadFromFile[discfg.Config](getConfigPath())
	panicOnErr(err)

	log.Info().Msg("Initing Texit API client")
	texit := initClient(cfg.Api.Address, cfg.Api.ApiKey)
	log.Info().Msg("Initing api service")
	apiGw := api_gateway.New(texit)
	apiSvc := api.NewService(apiGw)
	log.Info().Msg("Initing provider service")
	provSvc := provider.NewService(apiGw)
	log.Info().Msg("Initing tailnet service")
	tailSvc := tailnet.NewService(apiGw)
	log.Info().Msg("Initing node service")
	nodeSvc := node.NewService(apiGw, tailSvc, provSvc)
	log.Info().Msg("Initing handler")
	hdl := handler.New(apiSvc, nodeSvc, provSvc, tailSvc)

	log.Info().Msg("Creating new Tempest client...")
	client := tempest.NewClient(tempest.ClientOptions{
		PublicKey: cfg.Discord.PublicKey,
		Rest:      tempest.NewRest(cfg.Discord.Token),
	})

	log.Info().Msg("Initing Listener")
	lis := initListener(cfg.Server)

	log.Info().Msg("Initing Discord Bot")
	bot := discord.NewBot(lis, hdl, client)

	go func() {
		log.Info().Msg("Starting Bot")
		panicOnErr(bot.Start(ctx))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Shutting down bot")
	cancel()

	<-ctx.Done()

	log.Info().Msg("Exiting")
}

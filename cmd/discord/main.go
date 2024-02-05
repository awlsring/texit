package main

import (
	"context"
	"log/slog"
	"os"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord"
	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	api_gateway "github.com/awlsring/texit/internal/app/ui/adapters/secondary/gateway/api"
	"github.com/awlsring/texit/internal/app/ui/config"
	"github.com/awlsring/texit/internal/app/ui/core/service/api"
	"github.com/awlsring/texit/internal/pkg/logger"
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = logger.InitContextLogger(ctx, zerolog.DebugLevel)
	log = logger.FromContext(ctx)
	log.Info().Msg("Initializing")

	log.Info().Msg("Loading config")
	cfg, err := config.LoadFromFile[discfg.Config](getConfigPath())
	panicOnErr(err)

	texit := initClient(cfg.Api.Address, cfg.Api.ApiKey)
	apiGw := api_gateway.New(texit)
	svc := api.NewService(apiGw)
	hdl := handler.New(svc)

	log.Info().Msg("Creating new Tempest client...")
	client := tempest.NewClient(tempest.ClientOptions{
		PublicKey: cfg.Discord.PublicKey,
		Rest:      tempest.NewRest(cfg.Discord.Token),
	})

	log.Info().Msg("Initing Discord Bot")
	bot := discord.New(hdl, client)
	panicOnErr(bot.Initialize())

	log.Info().Msg("Syncing local commands")
	err = client.SyncCommands(nil, nil, false)
	if err != nil {
		slog.Error("failed to sync local commands storage with Discord API", err)
	}

	log.Info().Msg("Starting Discord Bot")
	if err := client.ListenAndServe("/", cfg.Server.Address); err != nil {
		slog.Error("something went terribly wrong", err)
	}
}

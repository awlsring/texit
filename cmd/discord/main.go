package main

import (
	"context"
	"os"
	"os/signal"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/callback"
	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	pending_execution "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/execution"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	api_gateway "github.com/awlsring/texit/internal/app/ui/adapters/secondary/gateway/api"
	"github.com/awlsring/texit/internal/app/ui/core/service/api"
	"github.com/awlsring/texit/internal/app/ui/core/service/node"
	"github.com/awlsring/texit/internal/app/ui/core/service/provider"
	"github.com/awlsring/texit/internal/app/ui/core/service/tailnet"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/config"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/mqtt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	configEnvVar          = "DISCORD_CONFIG_PATH"
	defaultConfigLocation = "/etc/texit_discord/config.yaml"
)

func getConfigPath() string {
	path := os.Getenv(configEnvVar)
	if path == "" {
		return defaultConfigLocation
	}
	return path
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Info().Msg("Initializing")

	log.Info().Msg("Loading config")
	cfg, err := config.LoadFromFile[discfg.Config](getConfigPath())
	appinit.PanicOnErr(err)
	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	zerolog.SetGlobalLevel(lvl)
	log := logger.InitLogger(lvl)
	appinit.PanicOnErr(err)

	log.Info().Msg("Initing Texit API client")
	texit := discord.LoadTexitClient(cfg.Api.Address, cfg.Api.ApiKey)

	log.Info().Msg("Initing api service")
	apiGw := api_gateway.New(texit)
	apiSvc := api.NewService(apiGw)

	log.Info().Msg("Initing provider service")
	provSvc := provider.NewService(apiGw)

	log.Info().Msg("Initing tailnet service")
	tailSvc := tailnet.NewService(apiGw)

	log.Info().Msg("Initing node service")
	nodeSvc := node.NewService(apiGw, tailSvc, provSvc)

	log.Info().Msg("Initing tracker")
	tracker := pending_execution.NewInMemoryTracker()

	log.Info().Msg("Initing handler")
	hdl := handler.New(apiSvc, nodeSvc, provSvc, tailSvc, tracker)

	log.Info().Msg("Creating new Tempest client...")
	client := tempest.NewClient(tempest.ClientOptions{
		PublicKey: cfg.Discord.PublicKey,
		Rest:      tempest.NewRest(cfg.Discord.Token),
	})

	log.Info().Msg("Initing Listener")
	lis := discord.LoadListener(cfg.Server)

	log.Info().Msg("loading authorized snowflakes")
	authorized := []tempest.Snowflake{}
	for _, id := range cfg.Discord.Authorized {
		s, err := tempest.StringToSnowflake(id)
		appinit.PanicOnErr(err)
		authorized = append(authorized, s)
	}

	var guilds []tempest.Snowflake
	guilds = nil
	if len(cfg.Discord.GuildIds) > 0 {
		for _, id := range cfg.Discord.GuildIds {
			s, err := tempest.StringToSnowflake(id)
			appinit.PanicOnErr(err)
			guilds = append(guilds, s)
		}
	}

	log.Info().Msg("Initing Callback Handler")
	lisHdl := callback.NewCallbackHandler(client, tracker)

	lsn, err := mqtt.NewListener(cfg.Notification.Broker, lisHdl, mqtt.WithLogLevel(zerolog.DebugLevel))
	appinit.PanicOnErr(err)

	log.Info().Msg("Initing Discord Bot")
	bot := discord.NewBot(hdl, client, discord.WithAuthorizedUsers(authorized), discord.WithGuilds(guilds), discord.WithLogLevel(lvl))

	go func() {
		log.Info().Msg("Starting Bot")
		appinit.PanicOnErr(bot.Serve(ctx, lis))
	}()

	go func() {
		log.Info().Msg("Starting MQTT Listener on topic " + cfg.Notification.Topic)
		appinit.PanicOnErr(lsn.Subscribe(ctx, cfg.Notification.Topic))
		log.Info().Msg("Subscribed to MQTT topic")
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Shutting down bot")
	cancel()

	<-ctx.Done()

	log.Info().Msg("Exiting")
}

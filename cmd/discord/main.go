package main

import (
	"context"
	"os"
	"os/signal"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/a-h/awsapigatewayv2handler"
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
	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/mqtt"
	"github.com/awlsring/texit/internal/pkg/runtime"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/rs/zerolog"
)

var (
	log         zerolog.Logger
	lvl         zerolog.Level
	cfg         *discfg.Config
	texitClient texit.Invoker
	tmpstClient *tempest.Client
	apiGw       gateway.Api
	apiSvc      service.Api
	provSvc     service.Provider
	tailSvc     service.Tailnet
	nodeSvc     service.Node
	tracker     pending_execution.Tracker
	bot         *discord.Bot
)

func main() {
	var err error
	log.Info().Msg("Initializing")

	log.Info().Msg("Loading config")
	cfg, err = discord.LoadConfig()
	appinit.PanicOnErr(err)
	err = cfg.Validate()
	appinit.PanicOnErr(err)
	lvl, err = zerolog.ParseLevel(cfg.LogLevel)
	zerolog.SetGlobalLevel(lvl)
	log = logger.InitLogger(lvl)
	appinit.PanicOnErr(err)

	log.Info().Msg("Initing Texit API client")
	texitClient = discord.LoadTexitClient(cfg.Api.Address, cfg.Api.ApiKey)

	log.Info().Msg("Initing api service")
	apiGw = api_gateway.New(texitClient)
	apiSvc = api.NewService(apiGw)

	log.Info().Msg("Initing provider service")
	provSvc = provider.NewService(apiGw)

	log.Info().Msg("Initing tailnet service")
	tailSvc = tailnet.NewService(apiGw)

	log.Info().Msg("Initing node service")
	nodeSvc = node.NewService(apiGw, tailSvc, provSvc)

	log.Info().Msg("Initing tracker")
	tracker, err = discord.LoadTracker(cfg.Tracker)
	appinit.PanicOnErr(err)

	log.Info().Msg("Initing handler")
	hdl := handler.New(apiSvc, nodeSvc, provSvc, tailSvc, tracker)

	log.Info().Msg("Creating new Tempest client...")
	tmpstClient = tempest.NewClient(tempest.ClientOptions{
		PublicKey: cfg.Discord.PublicKey,
		Rest:      tempest.NewRest(cfg.Discord.Token),
	})

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

	log.Info().Msg("Initing Discord Bot")
	bot = discord.NewBot(hdl, tmpstClient, discord.WithAuthorizedUsers(authorized), discord.WithGuilds(guilds), discord.WithLogLevel(lvl))

	if runtime.IsLambda() {
		startLambdaServer()
	} else {
		startServer()
	}
}

func startLambdaServer() {
	log.Info().Msg("Starting lambda bot")
	if cfg.Server.Address != ":443" {
		log.Warn().Msgf("Only :443 is supported as a server address, ignoring set address of %s", cfg.Server.Address)
	}

	if cfg.Tracker.Type == discfg.TrackerTypeInMemory {
		panic("Only tracker supported in lambda is DDB")
	}

	if cfg.Notification.Type != discfg.NotifierTypeSns {
		panic("Only notifier supported in lambda is SNS")
	}

	err := bot.Init()
	appinit.PanicOnErr(err)

	hdl := bot.HttpHandler()
	log.Info().Msg("Starting API Gateway V2 Handler")
	awsapigatewayv2handler.ListenAndServe(hdl)
}

func startServer() {
	log.Info().Msg("Prepping server launch")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info().Msg("Initing Listener")
	lis := discord.LoadListener(cfg.Server)

	go func() {
		log.Info().Msg("Starting Bot")
		appinit.PanicOnErr(bot.Serve(ctx, lis))
	}()

	if cfg.Notification.Type == discfg.NotifierTypeMqtt {
		log.Info().Msg("Initing Callback Handler")
		lisHdl := callback.NewCallbackHandler(tmpstClient, tracker)

		lsn, err := mqtt.NewListener(cfg.Notification.Broker, lisHdl, mqtt.WithLogLevel(zerolog.DebugLevel))
		appinit.PanicOnErr(err)
		go func() {
			log.Info().Msg("Starting MQTT Listener on topic " + cfg.Notification.Topic)
			appinit.PanicOnErr(lsn.Subscribe(ctx, cfg.Notification.Topic))
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Shutting down bot")
	cancel()

	<-ctx.Done()

	log.Info().Msg("Exiting")
}

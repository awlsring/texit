package main

import (
	"context"
	"os"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/a-h/awsapigatewayv2handler"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord"
	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	pending_execution "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/execution"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	api_gateway "github.com/awlsring/texit/internal/app/ui/adapters/secondary/gateway/api"
	"github.com/awlsring/texit/internal/app/ui/core/service/api"
	"github.com/awlsring/texit/internal/app/ui/core/service/node"
	"github.com/awlsring/texit/internal/app/ui/core/service/provider"
	"github.com/awlsring/texit/internal/app/ui/core/service/tailnet"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/clients"
	cconfig "github.com/awlsring/texit/internal/pkg/config"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func texitEndpoint(cfg *discfg.Config) string {
	addr := os.Getenv("TEXIT_ENDPOINT")
	if addr == "" {
		if cfg.Api.Address == "" {
			panic("TEXIT_ENDPOINT environment variable not set and not set in config")
		}
		addr = cfg.Api.Address
	}
	return addr
}

func main() {
	log.Info().Msg("Starting bot")
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background())
	appinit.PanicOnErr(err)

	log.Info().Msg("Loading from S3 config")
	s3Client := s3.NewFromConfig(awsCfg)
	cfg, err := cconfig.LoadFromS3[discfg.Config](s3Client, os.Getenv("CONFIG_BUCKET"), os.Getenv("CONFIG_OBJECT"))
	appinit.PanicOnErr(err)
	if cfg.Server.Address != ":443" {
		panic("Only :443 is supported as a server address")
	}
	ddbClient := dynamodb.NewFromConfig(awsCfg)

	log.Info().Msg("Initing Texit API client")
	texit, err := clients.CreateTexitClient(texitEndpoint(cfg), cfg.Api.ApiKey)
	appinit.PanicOnErr(err)

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
	tracker := pending_execution.NewDdbTracker("TrackedExecutions", ddbClient)

	log.Info().Msg("Initing handler")
	hdl := handler.New(apiSvc, nodeSvc, provSvc, tailSvc, tracker)

	log.Info().Msg("Creating new Tempest client...")
	client := tempest.NewClient(tempest.ClientOptions{
		PublicKey: cfg.Discord.PublicKey,
		Rest:      tempest.NewRest(cfg.Discord.Token),
	})

	log.Info().Msg("loading authorized snowflakes")
	authorized, err := cfg.Discord.AuthorizedAsSnowflakes()
	appinit.PanicOnErr(err)

	guilds, err := cfg.Discord.AuthorizedGuildsAsSnowflakes()
	appinit.PanicOnErr(err)

	log.Info().Msg("Initing Discord Bot")
	bot := discord.NewBot(hdl, client, discord.WithAuthorizedUsers(authorized), discord.WithGuilds(guilds), discord.WithLogLevel(zerolog.DebugLevel))

	appinit.PanicOnErr(bot.Init())

	awsapigatewayv2handler.ListenAndServe(bot.HttpHandler())
}

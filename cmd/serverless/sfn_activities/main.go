package main

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/sfn_activities"
	dynamo_execution_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/execution/dynamo"
	dynamo_node_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/node/dynamo"
	"github.com/awlsring/texit/internal/app/api/core/service/activity"
	"github.com/awlsring/texit/internal/app/api/core/service/notification"
	"github.com/awlsring/texit/internal/app/api/setup"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting AWS lambda activity handler")
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Error loading AWS config")
		appinit.PanicOnErr(err)
	}
	cfg, err := setup.LoadConfig()
	appinit.PanicOnErr(err)

	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	log := logger.InitLogger(lvl)
	log.Info().Msgf("Setting log level to %s", lvl.String())
	zerolog.SetGlobalLevel(lvl)
	appinit.PanicOnErr(err)

	ddb := dynamodb.NewFromConfig(awsCfg)
	nodeRepo := dynamo_node_repository.New("TexitNodes", ddb)
	execRepo := dynamo_execution_repository.New("TexitExecutions", ddb)

	log.Info().Msg("Initializing provider gateways")
	providerGateways := setup.LoadProviderGateways(cfg.Providers)

	log.Info().Msg("Initializing tailnet gateways")
	tailnetGateways := setup.LoadTailnetGateways(cfg.Tailnets)

	log.Info().Msg("Creating activity service")
	actSvc := activity.NewService(tailnetGateways, providerGateways, nodeRepo, execRepo)

	log.Info().Msg("Initializing notifier gateways")
	notifiers := setup.LoadNotifiers(cfg.Notifiers)

	log.Info().Msg("Initializing notification service")
	notSvc, err := notification.NewService(notifiers)
	appinit.PanicOnErr(err)

	log.Info().Msg("Creating activity handler")
	act := sfn_activities.New(notSvc, actSvc)

	log.Info().Msg("Starting lambda handler")
	lambda.Start(act.HandleRequest)
}

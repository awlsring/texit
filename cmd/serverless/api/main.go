package main

import (
	"context"
	"io"
	"os"

	"github.com/a-h/awsapigatewayv2handler"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/auth"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/handler"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/middleware"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/smithy_errors"
	"github.com/awlsring/texit/internal/app/api/setup"

	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/core/service/node"
	"github.com/awlsring/texit/internal/app/api/core/service/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loadAppConfig(acfg aws.Config) *config.Config {
	client := s3.NewFromConfig(acfg)
	bucketName := os.Getenv("CONFIG_BUCKET")
	if bucketName == "" {
		panic("CONFIG_BUCKET environment variable not set")
	}
	resp, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    aws.String("config.yaml"),
	})
	panicOnErr(err)
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	panicOnErr(err)
	cfg, err := config.LoadFromData(bytes)
	panicOnErr(err)
	return cfg
}

func main() {
	log.Info().Msg("Starting server...")
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Error loading AWS config")
		panicOnErr(err)
	}

	cfg := loadAppConfig(awsCfg)
	if cfg.Server.Address != ":443" {
		panic("Only :443 is supported as a server address")
	}
	if cfg.Database.Engine == config.DatabaseEngineSqlite {
		panic("Sqlite not supported in lambda")
	}

	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	log = logger.InitLogger(lvl)
	log.Info().Msgf("Setting log level to %s", lvl.String())
	zerolog.SetGlobalLevel(lvl)
	panicOnErr(err)

	log.Info().Msg("Connecting to database")
	nodeRepo, execRepo := setup.LoadRepositories(cfg.Database)

	log.Info().Msg("Initializing provider gateways")
	providerGateways := setup.LoadProviderGateways(cfg.Providers)

	log.Info().Msg("Initializing workflow gateway")
	workGw := setup.LoadStepFunctionsWorkflowGateway(awsCfg)

	log.Info().Msg("Initializing workflow service")
	workflowSvc := workflow.NewService(nodeRepo, execRepo, workGw)

	log.Info().Msg("Initializing provider service")
	providerSvc := setup.LoadProviderService(cfg.Providers)

	tailnetSvc := setup.LoadTailnetService(cfg.Tailnets)

	log.Info().Msg("Initializing node service")
	nodeSvc := node.NewService(nodeRepo, workflowSvc, providerGateways)

	log.Info().Msg("Froming ogen handler")
	hdl := handler.New(nodeSvc, workflowSvc, providerSvc, tailnetSvc)

	log.Info().Msg("Initializing security handler")
	sec := auth.NewSecurityHandler([]string{cfg.Server.APIKey})
	opts := []texit.ServerOption{
		texit.WithMiddleware(middleware.LoggingMiddleware(lvl)),
		texit.WithNotFound(smithy_errors.UnknownOperationHandler),
		texit.WithErrorHandler(smithy_errors.ResponseHandlerWithLogger(lvl)),
	}
	srv, err := texit.NewServer(hdl, sec, opts...)
	if err != nil {
		log.Error().Err(err).Msg("Error creating server")
		panicOnErr(err)
	}

	awsapigatewayv2handler.ListenAndServe(srv)
}

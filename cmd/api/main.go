package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/mem_workflow"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/auth"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/handler"
	local_workflow "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/workflow/local"
	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/core/service/activity"
	"github.com/awlsring/texit/internal/app/api/core/service/node"
	"github.com/awlsring/texit/internal/app/api/core/service/notification"
	workflowSvc "github.com/awlsring/texit/internal/app/api/core/service/workflow"
	"github.com/awlsring/texit/internal/app/api/setup"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

const (
	configEnvVar          = "CONFIG_PATH"
	defaultConfigLocation = "/etc/texit/config.yaml"
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
	cfg, err := config.LoadFromFile(getConfigPath())
	appinit.PanicOnErr(err)

	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	log = logger.InitLogger(lvl)
	log.Info().Msgf("Setting log level to %s", lvl.String())
	zerolog.SetGlobalLevel(lvl)
	appinit.PanicOnErr(err)

	log.Info().Msg("Connecting to database")
	nodeRepo, excRepo := setup.LoadRepositories(cfg.Database)

	log.Info().Msg("Initializing provider gateways")
	providerGateways := setup.LoadProviderGateways(cfg.Providers)

	log.Info().Msg("Initializing tailnet gateways")
	tailnetGateways := setup.LoadTailnetGateways(cfg.Tailnets)

	log.Info().Msg("Initializing activity service")
	activitySvc := activity.NewService(tailnetGateways, providerGateways, nodeRepo, excRepo)

	workChan := make(chan workflow.ExecutionInput)

	log.Info().Msg("Initializing workflow gateway")
	workGw := local_workflow.New(workChan)

	log.Info().Msg("Initializing workflow service")
	workflowSvc := workflowSvc.NewService(nodeRepo, excRepo, workGw)

	log.Info().Msg("Initializing provider service")
	providerSvc := setup.LoadProviderService(cfg.Providers)

	tailnetSvc := setup.LoadTailnetService(cfg.Tailnets)

	log.Info().Msg("Initializing notifier gateways")
	notifiers := setup.LoadNotifiers(cfg.Notifiers)

	log.Info().Msg("Initializing notification service")
	notSvc := notification.NewService(notifiers)

	log.Info().Msg("Initializing node service")
	nodeSvc := node.NewService(nodeRepo, workflowSvc, providerGateways)

	log.Info().Msg("Froming ogen handler")
	hdl := handler.New(nodeSvc, workflowSvc, providerSvc, tailnetSvc)

	log.Info().Msg("Initializing security handler")
	sec := auth.NewSecurityHandler([]string{cfg.Server.APIKey})

	log.Info().Msg("Creating ogen server")
	srv, err := ogen.NewServer(hdl, ogen.WithSecurityHandler(sec), ogen.WithLogLevel(lvl))
	appinit.PanicOnErr(err)

	log.Info().Msg("Initializing net listener")
	lis := setup.LoadListener(cfg.Server)

	log.Info().Msg("Initializing workflow worker")
	worker := mem_workflow.NewWorker(activitySvc, notSvc, workChan, mem_workflow.WithLogLevel(lvl))

	log.Info().Msg("Starting server")
	go func() {
		appinit.PanicOnErr(srv.Serve(ctx, lis))
	}()

	log.Info().Msg("Starting worker")
	go func() {
		appinit.PanicOnErr(worker.Start(ctx))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Shutting down server")
	cancel()

	log.Info().Msg("Waiting for server to shutdown")
	<-ctx.Done()

	appinit.PanicOnErr(worker.Close(context.Background()))
	nodeRepo.Close()
	excRepo.Close()

	log.Info().Msg("Exiting")
}

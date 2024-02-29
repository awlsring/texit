package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/a-h/awsapigatewayv2handler"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/mem_workflow"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/auth"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/handler"
	local_workflow "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/workflow/local"
	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/core/service/activity"
	"github.com/awlsring/texit/internal/app/api/core/service/node"
	"github.com/awlsring/texit/internal/app/api/core/service/notification"
	wrkSvc "github.com/awlsring/texit/internal/app/api/core/service/workflow"
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/app/api/setup"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/runtime"
	"github.com/rs/zerolog"
)

var (
	log              zerolog.Logger
	lvl              zerolog.Level
	cfg              *config.Config
	nodeRepo         repository.Node
	execRepo         repository.Execution
	providerGateways map[string]gateway.Platform
	tailnetGateways  map[string]gateway.Tailnet
	notifierGateways map[string]gateway.Notification
	workflowGateway  gateway.Workflow
	activitySvc      service.Activity
	nodeSvc          service.Node
	workflowSvc      service.Workflow
	providerSvc      service.Provider
	tailnetSvc       service.Tailnet
	notSvc           service.Notification
	srv              *ogen.Server
)

func main() {
	log.Info().Msg("Initializing")
	var err error

	log.Info().Msg("Loading config")
	cfg, err = setup.LoadConfig()
	appinit.PanicOnErr(err)

	lvl, err = zerolog.ParseLevel(cfg.LogLevel)
	log = logger.InitLogger(lvl)
	log.Info().Msgf("Setting log level to %s", lvl.String())
	zerolog.SetGlobalLevel(lvl)
	appinit.PanicOnErr(err)

	log.Info().Msg("Connecting to database")
	nodeRepo, execRepo = setup.LoadRepositories(cfg.Database)

	log.Info().Msg("Initializing provider gateways")
	providerGateways = setup.LoadProviderGateways(cfg.Providers)

	log.Info().Msg("Initializing tailnet gateways")
	tailnetGateways = setup.LoadTailnetGateways(cfg.Tailnets)

	log.Info().Msg("Initializing activity service")
	activitySvc = activity.NewService(tailnetGateways, providerGateways, nodeRepo, execRepo)

	log.Info().Msg("Initializing workflow gateway")
	workflowGateway = setup.LoadWorkflowEngine(cfg.Workflow)

	log.Info().Msg("Initializing workflow service")
	workflowSvc = wrkSvc.NewService(nodeRepo, execRepo, workflowGateway)

	log.Info().Msg("Initializing provider service")
	providerSvc = setup.LoadProviderService(cfg.Providers)

	tailnetSvc = setup.LoadTailnetService(cfg.Tailnets)

	log.Info().Msg("Initializing notifier gateways")
	notifierGateways = setup.LoadNotifiers(cfg.Notifiers)

	log.Info().Msg("Initializing notification service")
	notSvc, err = notification.NewService(notifierGateways)
	appinit.PanicOnErr(err)

	log.Info().Msg("Initializing node service")
	nodeSvc = node.NewService(nodeRepo, workflowSvc, providerGateways)

	log.Info().Msg("Froming ogen handler")
	hdl := handler.New(nodeSvc, workflowSvc, providerSvc, tailnetSvc, notSvc)

	log.Info().Msg("Initializing security handler")
	sec := auth.NewSecurityHandler([]string{cfg.Server.APIKey})

	log.Info().Msg("Creating ogen server")
	srv, err = ogen.NewServer(hdl, ogen.WithSecurityHandler(sec), ogen.WithLogLevel(lvl))
	appinit.PanicOnErr(err)

	if runtime.IsLambda() {
		startLambdaServer()
	} else {
		startServer()
	}
}

func startLambdaServer() {
	log.Info().Msg("Starting lambda server")

	// validate non lambda things wont happen here
	if cfg.Server.Address != ":443" {
		log.Warn().Msgf("Only :443 is supported as a server address, ignoring set address of %s", cfg.Server.Address)
	}

	if cfg.Database.Engine == config.DatabaseEngineSqlite {
		panic("Sqlite not supported in lambda")
	}

	if cfg.Workflow.Type == "local" {
		panic("Local workflow not supported in lambda")
	}

	hdl := srv.HttpHandler()
	awsapigatewayv2handler.ListenAndServe(hdl)
}

func startServer() {
	log.Info().Msg("Prepping server launch")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Info().Msg("Initializing net listener")
	lis := setup.LoadListener(cfg.Server)

	log.Info().Msg("Starting server")
	go func() {
		appinit.PanicOnErr(srv.Serve(ctx, lis))
	}()

	if cfg.Workflow.Type == "local" {
		log.Info().Msg("Starting worker")
		localWorkGw := workflowGateway.(*local_workflow.LocalWorkflow)
		worker := mem_workflow.NewWorker(activitySvc, notSvc, localWorkGw.Channel(), mem_workflow.WithLogLevel(lvl))
		defer worker.Close(context.Background())
		go func() {
			log.Info().Msg("Initializing workflow worker")
			appinit.PanicOnErr(worker.Start(ctx))
		}()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Shutting down server")
	cancel()

	log.Info().Msg("Waiting for server to shutdown")
	<-ctx.Done()

	nodeRepo.Close()
	execRepo.Close()

	defer log.Info().Msg("Exiting")
}

package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/handler"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ecs"
	tailscale_gateway "github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/secondary/gateway/tailnet/tailscale"
	sqlite_node_repository "github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/secondary/repository/sqlite"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/config"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/service/node"
	provSvc "github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/service/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/service/workflow"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"github.com/tailscale/tailscale-client-go/tailscale"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func loadProviderGateways(providers []config.ProviderConfig) (map[string]gateway.Platform, error) {
	gateways := make(map[string]gateway.Platform)
	for _, provider := range providers {
		switch provider.Type {
		case "aws-ecs":
			p := platform_aws_ecs.New(provider.AccessKey, provider.SecretKey)
			gateways[provider.Name] = p
		default:
			return nil, nil
		}
	}
	return gateways, nil
}

func initProviderService(providers []config.ProviderConfig) service.Provider {
	provs := []*provider.Provider{}
	for _, p := range providers {
		name, err := provider.IdentifierFromString(p.Name)
		panicOnErr(err)
		typ, err := provider.TypeFromString(p.Type)
		panicOnErr(err)
		provs = append(provs, &provider.Provider{
			Name:     name,
			Platform: typ,
			Default:  p.Default,
		})
	}
	svc, err := provSvc.NewService(provs)
	panicOnErr(err)
	return svc
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = logger.InitContextLogger(ctx, zerolog.DebugLevel)
	log := logger.FromContext(ctx)
	log.Info().Msg("Initializing")

	log.Info().Msg("Loading config")
	cfg, err := config.LoadFromFile("config.yaml")
	panicOnErr(err)
	// log.Debug().Interface("config", cfg).Msg("Loaded config")

	log.Info().Msg("Connecting to database")
	db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	panicOnErr(err)
	nodeRepo := sqlite_node_repository.New(db)
	err = nodeRepo.Init(ctx)
	panicOnErr(err)

	log.Info().Msg("Initializing tailscale client")
	ts, err := tailscale.NewClient(cfg.Tailscale.ApiKey, cfg.Tailscale.Tailnet)
	panicOnErr(err)
	log.Info().Msg("Initializing tailscale gateway")
	tailnetGateway := tailscale_gateway.New(ts)

	log.Info().Msg("Initializing provider gateways")
	providerGateways, err := loadProviderGateways(cfg.Providers)
	panicOnErr(err)

	log.Info().Msg("Initializing workflow service")
	workflowSvc := workflow.NewService(nodeRepo, tailnetGateway, providerGateways)

	log.Info().Msg("Initializing provider service")
	providerSvc := initProviderService(cfg.Providers)

	log.Info().Msg("Initializing node service")
	nodeSvc := node.NewService(nodeRepo, workflowSvc, providerGateways)

	log.Info().Msg("Froming gRPC handler")
	hdl := handler.New(nodeSvc, workflowSvc, providerSvc)

	log.Info().Msg("Creating gRPC server")
	srv, err := grpc.NewServer(hdl, grpc.WithLogLevel(zerolog.DebugLevel))
	panicOnErr(err)

	log.Info().Msg("Starting gRPC server")
	go func() {
		panicOnErr(srv.Start(ctx))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c
	log.Info().Msg("Shutting down server")
	cancel()

	log.Info().Msg("Waiting for server to shutdown")
	<-ctx.Done()

	err = db.Close()
	panicOnErr(err)

	log.Info().Msg("Exiting")
}

package main

import (
	"context"
	"net/url"
	"os"
	"os/signal"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/handler"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ecs"
	headscale_v0_22_3_gateway "github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/secondary/gateway/tailnet/headscale/v0.22.3"
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
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

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
		typ, err := provider.TypeFromString(p.Type.String())
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

func initTailnetGateway(cfg config.TailnetConfig) gateway.Tailnet {
	switch cfg.Type {
	case config.TailnetTypeTailscale:
		return initTailscaleGateway(cfg)
	default:
		panic("invalid tailnet type")
	}
}

func initTailscaleGateway(cfg config.TailnetConfig) gateway.Tailnet {
	log.Info().Msg("Initializing tailscale client")
	ts, err := tailscale.NewClient(cfg.ApiKey, cfg.Tailnet)
	panicOnErr(err)
	log.Info().Msg("Initializing tailscale gateway")
	return tailscale_gateway.New(ts)
}

func initHeadscaleGateway(cfg config.TailnetConfig) gateway.Tailnet {
	u, err := url.Parse(cfg.Tailnet)
	panicOnErr(err)
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(cfg.ApiKey)

	client := client.New(transport, strfmt.Default)

	return headscale_v0_22_3_gateway.New(cfg.User, client.HeadscaleService)
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

	log.Info().Msg("Connecting to database")
	db, err := sqlx.Connect("sqlite3", "__deleteme.db")
	panicOnErr(err)
	nodeRepo := sqlite_node_repository.New(db)
	err = nodeRepo.Init(ctx)
	panicOnErr(err)

	log.Info().Msg("Initializing tailnet gateway")
	tailnetGateway := initTailnetGateway(cfg.Tailnet)

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

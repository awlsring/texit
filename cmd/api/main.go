package main

import (
	"context"
	"net"
	"net/url"
	"os"
	"os/signal"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/auth"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/handler"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ec2"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ecs"
	headscale_v0_22_3_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/tailnet/headscale/v0.22.3"
	tailscale_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/tailnet/tailscale"
	sqlite_execution_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/execution/sqlite"
	sqlite_node_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/node/sqlite"
	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/core/service/node"
	provSvc "github.com/awlsring/texit/internal/app/api/core/service/provider"
	tailnetSvc "github.com/awlsring/texit/internal/app/api/core/service/tailnet"
	"github.com/awlsring/texit/internal/app/api/core/service/workflow"
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/tsn"
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/tailscale/tailscale-client-go/tailscale"
	_ "modernc.org/sqlite"
)

var log zerolog.Logger

const (
	configEnvVar          = "CONFIG_PATH"
	defaultConfigLocation = "/etc/texit/config.yaml"
)

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

func initProviderGateways(providers []*config.ProviderConfig) map[string]gateway.Platform {
	gateways := make(map[string]gateway.Platform)
	for _, provider := range providers {
		switch provider.Type {
		case "aws-ecs":
			p := platform_aws_ecs.New(provider.AccessKey, provider.SecretKey)
			gateways[provider.Name] = p
		case "aws-ec2":
			p := platform_aws_ec2.New(provider.AccessKey, provider.SecretKey)
			gateways[provider.Name] = p
		default:
			return nil
		}
	}
	return gateways
}

func initTailnetGateways(cfg []*config.TailnetConfig) map[string]gateway.Tailnet {
	gateways := make(map[string]gateway.Tailnet)
	for _, t := range cfg {
		_, ok := gateways[t.Tailnet]
		if ok {
			panic("duplicate tailnet specified in config file")
		}
		switch t.Type {
		case config.TailnetTypeTailscale:
			gateways[t.Tailnet] = initTailscaleGateway(t)
		case config.TailnetTypeHeadscale:
			gateways[t.Tailnet] = initHeadscaleGateway(t)
		default:
			return nil
		}
	}
	return gateways
}

func initProviderService(providers []*config.ProviderConfig) service.Provider {
	provs := []*provider.Provider{}
	for _, p := range providers {
		name, err := provider.IdentifierFromString(p.Name)
		panicOnErr(err)
		typ, err := provider.TypeFromString(p.Type.String())
		panicOnErr(err)
		provs = append(provs, &provider.Provider{
			Name:     name,
			Platform: typ,
		})
	}
	svc := provSvc.NewService(provs)
	return svc
}

func initTailnetService(tailnets []*config.TailnetConfig) service.Tailnet {
	provs := []*tailnet.Tailnet{}
	for _, t := range tailnets {
		name, err := tailnet.IdentifierFromString(t.Tailnet)
		panicOnErr(err)
		typ, err := tailnet.TypeFromString(t.Type.String())
		panicOnErr(err)
		cs, err := tailnet.ControlServerFromString(t.ControlServer)
		panicOnErr(err)
		provs = append(provs, &tailnet.Tailnet{
			Name:          name,
			Type:          typ,
			ControlServer: cs,
		})
	}
	svc := tailnetSvc.NewService(provs)
	return svc
}

func initTailscaleGateway(cfg *config.TailnetConfig) gateway.Tailnet {
	log.Info().Msg("Initializing tailscale client")
	ts, err := tailscale.NewClient(cfg.ApiKey, cfg.Tailnet)
	panicOnErr(err)
	log.Info().Msg("Initializing tailscale gateway")
	return tailscale_gateway.New(cfg.User, ts)
}

func initHeadscaleGateway(cfg *config.TailnetConfig) gateway.Tailnet {
	u, err := url.Parse(cfg.ControlServer)
	panicOnErr(err)
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(cfg.ApiKey)

	client := client.New(transport, strfmt.Default)

	return headscale_v0_22_3_gateway.New(cfg.User, client.HeadscaleService)
}

func initListener(cfg *config.ServerConfig) net.Listener {
	if cfg.Tailnet != nil {
		l, err := tsn.ListenerFromConfig(*cfg.Tailnet, cfg.Address, tsn.WithStandardLoggingFunc(log))
		panicOnErr(err)
		return l
	}
	log.Info().Msg("Creating normal net listener")
	l, err := net.Listen("tcp", cfg.Address)
	panicOnErr(err)
	return l
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx = logger.InitContextLogger(ctx, zerolog.DebugLevel)
	log = logger.FromContext(ctx)
	log.Info().Msg("Initializing")

	log.Info().Msg("Loading config")
	cfg, err := config.LoadFromFile(getConfigPath())
	panicOnErr(err)

	log.Info().Msg("Connecting to database")
	db, err := sqlx.Connect("sqlite", cfg.Database.Location)
	panicOnErr(err)
	nodeRepo := sqlite_node_repository.New(db)
	err = nodeRepo.Init(ctx)
	panicOnErr(err)

	excRepo := sqlite_execution_repository.New(db)
	err = excRepo.Init(ctx)
	panicOnErr(err)

	log.Info().Msg("Initializing tailnet gateway")
	tailnetGateways := initTailnetGateways(cfg.Tailnets)

	log.Info().Msg("Initializing provider gateways")
	providerGateways := initProviderGateways(cfg.Providers)

	log.Info().Msg("Initializing workflow service")
	workflowSvc := workflow.NewService(nodeRepo, excRepo, tailnetGateways, providerGateways)

	log.Info().Msg("Initializing provider service")
	providerSvc := initProviderService(cfg.Providers)

	tailnetSvc := initTailnetService(cfg.Tailnets)

	log.Info().Msg("Initializing node service")
	nodeSvc := node.NewService(nodeRepo, workflowSvc, providerGateways)

	log.Info().Msg("Froming ogen handler")
	hdl := handler.New(nodeSvc, workflowSvc, providerSvc, tailnetSvc)

	log.Info().Msg("Initializing net listener")
	lis := initListener(cfg.Server)

	log.Info().Msg("Initializing security handler")
	sec := auth.NewSecurityHandler([]string{cfg.Server.APIKey})

	log.Info().Msg("Creating ogen server")
	srv := ogen.NewServer(lis, hdl, ogen.WithSecurityHandler(sec), ogen.WithLogLevel(zerolog.DebugLevel))

	log.Info().Msg("Starting server")
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

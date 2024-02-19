package main

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/a-h/awsapigatewayv2handler"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/auth"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/handler"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/middleware"
	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/smithy_errors"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ec2"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ecs"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_linode"
	headscale_v0_22_3_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/tailnet/headscale/v0.22.3"
	tailscale_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/tailnet/tailscale"
	step_functions_workflow "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/workflow/step_functions"
	dynamo_execution_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/execution/dynamo"
	dynamo_node_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/node/dynamo"
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
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/client"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/linode/linodego"
	"github.com/rs/zerolog"
	"github.com/tailscale/tailscale-client-go/tailscale"
	"golang.org/x/oauth2"
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
		case "linode":
			tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: provider.ApiKey})
			oauth2Client := &http.Client{
				Transport: &oauth2.Transport{
					Source: tokenSource,
				},
			}
			client := linodego.NewClient(oauth2Client)
			p := platform_linode.New(&client)
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
	return tailscale_gateway.New(ts)
}

func initHeadscaleGateway(cfg *config.TailnetConfig) gateway.Tailnet {
	u, err := url.Parse(cfg.ControlServer)
	panicOnErr(err)
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(cfg.ApiKey)

	client := client.New(transport, strfmt.Default)

	return headscale_v0_22_3_gateway.New(cfg.User, client.HeadscaleService)
}

func initWorkflowGateway(cfg aws.Config) gateway.Workflow {
	provArn := os.Getenv("PROVISION_NODE_WORKFLOW_ARN")
	deprovArn := os.Getenv("DEPROVISION_NODE_WORKFLOW_ARN")
	if provArn == "" || deprovArn == "" {
		panic("PROVISION_NODE_WORKFLOW_ARN and DEPROVISION_NODE_WORKFLOW_ARN must be set")
	}
	states := sfn.NewFromConfig(cfg)
	return step_functions_workflow.New(provArn, deprovArn, states)
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
	if cfg.Database.Engine != config.DatabaseEngineDynamoDb {
		panic("Only DynamoDB is supported as a database engine")
	}

	lvl, err := zerolog.ParseLevel(cfg.LogLevel)
	log = logger.InitLogger(lvl)
	log.Info().Msgf("Setting log level to %s", lvl.String())
	zerolog.SetGlobalLevel(lvl)
	panicOnErr(err)

	ddb := dynamodb.NewFromConfig(awsCfg)

	nodeRepo := dynamo_node_repository.New("TexitNodes", ddb)
	execRepo := dynamo_execution_repository.New("TexitExecutions", ddb)

	log.Info().Msg("Initializing provider gateways")
	providerGateways := initProviderGateways(cfg.Providers)

	log.Info().Msg("Initializing workflow gateway")
	workGw := initWorkflowGateway(awsCfg)

	log.Info().Msg("Initializing workflow service")
	workflowSvc := workflow.NewService(nodeRepo, execRepo, workGw)

	log.Info().Msg("Initializing provider service")
	providerSvc := initProviderService(cfg.Providers)

	tailnetSvc := initTailnetService(cfg.Tailnets)

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

package setup

import (
	"context"
	"net/http"
	"net/url"
	"os"

	mqtt_notification_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/notification/mqtt"
	sns_notification_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/notification/sns"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ec2"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_aws_ecs"
	"github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/platform/platform_linode"
	headscale_v0_22_3_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/tailnet/headscale/v0.22.3"
	tailscale_gateway "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/tailnet/tailscale"
	step_functions_workflow "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/workflow/step_functions"
	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/pkg/gen/headscale/v0.22.3/client"
	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	_ "github.com/lib/pq"
	"github.com/linode/linodego"
	"github.com/rs/zerolog/log"
	"github.com/tailscale/tailscale-client-go/tailscale"
	"golang.org/x/oauth2"
	_ "modernc.org/sqlite"
)

func LoadProviderGateways(providers []*config.ProviderConfig) map[string]gateway.Platform {
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

func LoadTailnetGateways(cfg []*config.TailnetConfig) map[string]gateway.Tailnet {
	gateways := make(map[string]gateway.Tailnet)
	for _, t := range cfg {
		_, ok := gateways[t.Tailnet]
		if ok {
			panic("duplicate tailnet specified in config file")
		}
		switch t.Type {
		case config.TailnetTypeTailscale:
			gateways[t.Tailnet] = LoadTailscaleGateway(t)
		case config.TailnetTypeHeadscale:
			gateways[t.Tailnet] = LoadHeadscaleGateway(t)
		default:
			return nil
		}
	}
	return gateways
}

func LoadTailscaleGateway(cfg *config.TailnetConfig) gateway.Tailnet {
	log.Info().Msg("Initializing tailscale client")
	ts, err := tailscale.NewClient(cfg.ApiKey, cfg.Tailnet)
	appinit.PanicOnErr(err)
	log.Info().Msg("Initializing tailscale gateway")
	return tailscale_gateway.New(ts)
}

func LoadHeadscaleGateway(cfg *config.TailnetConfig) gateway.Tailnet {
	u, err := url.Parse(cfg.ControlServer)
	appinit.PanicOnErr(err)
	transport := httptransport.New(u.Host, u.Path, []string{u.Scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(cfg.ApiKey)

	client := client.New(transport, strfmt.Default)

	return headscale_v0_22_3_gateway.New(cfg.User, client.HeadscaleService)
}

func LoadNotifiers(cfg []*config.NotifierConfig) []gateway.Notification {
	notifiers := make([]gateway.Notification, 0, len(cfg))
	for _, n := range cfg {
		switch n.Type {
		case config.NotifierTypeMqtt:
			opts := mqtt.NewClientOptions()
			opts.AddBroker(n.Broker)
			opts.SetClientID("texit")
			if n.Username != "" {
				opts.SetUsername(n.Username)
			}
			if n.Password != "" {
				opts.SetPassword(n.Password)
			}
			c := mqtt.NewClient(opts)
			if token := c.Connect(); token.Wait() && token.Error() != nil {
				appinit.PanicOnErr(token.Error())
			}
			notifiers = append(notifiers, mqtt_notification_gateway.New(n.Topic, c))
		case config.NotifierTypeSns:
			if n.AccessKey == "" || n.SecretKey == "" {
				panic("missing access key or secret key")
			}
			creds := credentials.NewStaticCredentialsProvider(n.AccessKey, n.SecretKey, "")
			cfg, err := awscfg.LoadDefaultConfig(context.TODO(),
				awscfg.WithRegion(n.Region),
				awscfg.WithCredentialsProvider(creds),
			)
			appinit.PanicOnErr(err)
			client := sns.NewFromConfig(cfg)
			notifiers = append(notifiers, sns_notification_gateway.New(n.Topic, client))
		default:
			panic("unknown notifier type")
		}
	}
	return notifiers
}

func LoadStepFunctionsWorkflowGateway(cfg aws.Config) gateway.Workflow {
	provArn := os.Getenv("PROVISION_NODE_WORKFLOW_ARN")
	deprovArn := os.Getenv("DEPROVISION_NODE_WORKFLOW_ARN")
	if provArn == "" || deprovArn == "" {
		panic("PROVISION_NODE_WORKFLOW_ARN and DEPROVISION_NODE_WORKFLOW_ARN must be set")
	}
	states := sfn.NewFromConfig(cfg)
	return step_functions_workflow.New(provArn, deprovArn, states)
}

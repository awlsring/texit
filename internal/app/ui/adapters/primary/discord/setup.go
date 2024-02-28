package discord

import (
	"context"
	"errors"
	"io"
	"net"
	"os"

	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	pending_execution "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/execution"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/config"
	cconfig "github.com/awlsring/texit/internal/pkg/config"
	"github.com/awlsring/texit/internal/pkg/runtime"
	"github.com/awlsring/texit/internal/pkg/tsn"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog/log"
)

const (
	configEnvVar          = "DISCORD_CONFIG_PATH"
	defaultConfigLocation = "/etc/texit_discord/config.yaml"
)

type Sec struct {
	key string
}

func (s Sec) SmithyAPIHttpApiKeyAuth(ctx context.Context, operationName string) (texit.SmithyAPIHttpApiKeyAuth, error) {
	return texit.SmithyAPIHttpApiKeyAuth{
		APIKey: s.key,
	}, nil
}

func getConfigPath() string {
	path := os.Getenv(configEnvVar)
	if path == "" {
		return defaultConfigLocation
	}
	return path
}

func LoadConfig() (*discfg.Config, error) {
	if runtime.IsLambda() {
		return loadConfigFromS3()
	}
	return config.LoadFromFile[discfg.Config](getConfigPath())
}

func loadConfigFromS3() (*discfg.Config, error) {
	acfg, err := awsconfig.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(acfg)
	bucketName := os.Getenv("CONFIG_BUCKET")
	if bucketName == "" {
		panic("CONFIG_BUCKET environment variable not set")
	}
	resp, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    aws.String("config.yaml"),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadFromData[discfg.Config](bytes)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadTexitClient(address string, key string) texit.Invoker {
	c, err := texit.NewClient(address, Sec{key: key})
	appinit.PanicOnErr(err)
	return c
}

func LoadListener(cfg discfg.ServerConfig) net.Listener {
	if cfg.Tailnet != nil {
		l, err := tsn.ListenerFromConfig(*cfg.Tailnet, cfg.Address)
		appinit.PanicOnErr(err)
		return l
	}
	log.Info().Msg("Creating net listener")
	l, err := net.Listen("tcp", cfg.Address)
	appinit.PanicOnErr(err)
	return l
}

func LoadTracker(cfg *discfg.TrackerConfig) (pending_execution.Tracker, error) {
	switch cfg.Type {
	case discfg.TrackerTypeInMemory:
		tracker := pending_execution.NewInMemoryTracker()
		return tracker, nil
	case discfg.TrackerTypeDynamoDb:
		awsCfg, err := cconfig.LoadAwsConfig(cfg.AccessKey, cfg.SecretKey, cfg.Region)
		if err != nil {
			return nil, err
		}
		ddbClient := dynamodb.NewFromConfig(awsCfg)
		tracker := pending_execution.NewDdbTracker("TrackedExecutions", ddbClient)
		return tracker, nil
	default:
		return nil, errors.New("unknown tracker type")
	}
}

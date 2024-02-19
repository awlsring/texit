package main

import (
	"context"
	"os"
	"sync"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/callback"
	discfg "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/config"
	pending_execution "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/execution"
	"github.com/awlsring/texit/internal/pkg/appinit"
	cconfig "github.com/awlsring/texit/internal/pkg/config"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var lvl zerolog.Level
var lisHdl *callback.CallbackHandler

func handleRecord(ctx context.Context, record events.SNSEventRecord) {
	log.Debug().Msg("Deserializing message")
	m, err := notification.DeserializeExecutionMessage([]byte(record.SNS.Message))
	if err != nil {
		log.Error().Err(err).Msg("Failed to deserialize message")
		return
	}

	lisHdl.Handle(ctx, m)
}

func handler(ctx context.Context, event events.SNSEvent) {
	ctx = logger.InitContextLogger(ctx, lvl)
	log := logger.FromContext(ctx)
	log.Debug().Interface("event", event).Msg("Handling event")

	var wg sync.WaitGroup
	for _, record := range event.Records {
		wg.Add(1)
		go func(record events.SNSEventRecord) {
			defer wg.Done()
			handleRecord(ctx, record)
		}(record)
	}

	wg.Wait()

	log.Debug().Msg("Done")
}

func main() {
	log.Info().Msg("Starting bot callback")
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background())
	appinit.PanicOnErr(err)

	log.Info().Msg("Loading from S3 config")
	s3Client := s3.NewFromConfig(awsCfg)
	ddbClient := dynamodb.NewFromConfig(awsCfg)
	cfg, err := cconfig.LoadFromS3[discfg.Config](s3Client, os.Getenv("CONFIG_BUCKET"), os.Getenv("CONFIG_OBJECT"))
	appinit.PanicOnErr(err)
	lvl, err = zerolog.ParseLevel(cfg.LogLevel)
	appinit.PanicOnErr(err)

	log.Info().Msg("Creating new Tempest client...")
	client := tempest.NewClient(tempest.ClientOptions{
		PublicKey: cfg.Discord.PublicKey,
		Rest:      tempest.NewRest(cfg.Discord.Token),
	})

	log.Info().Msg("Initing tracker")
	tracker := pending_execution.NewDdbTracker("TrackedExecutions", ddbClient)

	log.Info().Msg("Initing Callback Handler")
	lisHdl = callback.NewCallbackHandler(client, tracker)

	lambda.Start(handler)
}

package discord

import (
	"context"
	"net"
	"net/http"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type Bot struct {
	logLevel zerolog.Level
	tmpst    *tempest.Client
	hdl      *handler.Handler
	lis      net.Listener
}

func (b *Bot) SetLogLevel(lvl zerolog.Level) {
	b.logLevel = lvl
}

func (b *Bot) Handler() *handler.Handler {
	return b.hdl
}

func (b *Bot) Tempest() *tempest.Client {
	return b.tmpst
}

func NewBot(lis net.Listener, hdl *handler.Handler, tmpst *tempest.Client) *Bot {
	return &Bot{
		logLevel: zerolog.InfoLevel,
		lis:      lis,
		tmpst:    tmpst,
		hdl:      hdl,
	}
}

func (b *Bot) registerCommands() error {
	if err := b.tmpst.RegisterCommand(command.NewServerHealthCommand(b.logLevel, b.tmpst, b.hdl)); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewSelfHealthCheckCommand(b.hdl)); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeExecutionCommand(b.logLevel, b.tmpst, b.hdl)); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewProvisionNodeCommand(b.logLevel, b.tmpst, b.hdl)); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDeprovisionNodeCommand(b.logLevel, b.tmpst, b.hdl)); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewListNodesCommand(b.logLevel, b.tmpst, b.hdl)); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeNodeCommand(b.logLevel, b.tmpst, b.hdl)); err != nil {
		return err
	}

	return nil
}

func (b *Bot) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)

	if err := b.registerCommands(); err != nil {
		log.Error().Err(err).Msg("Failed to register commands")
		return err
	}

	guild, err := tempest.StringToSnowflake("948052547795574794")
	if err != nil {
		return err
	}
	if err := b.tmpst.SyncCommands([]tempest.Snowflake{guild}, nil, false); err != nil {
		return err
	}

	go func() {
		// take control of lifecycle so we can use our own serve methodology
		httpHdl, err := b.tmpst.Hijack()
		if err != nil {
			log.Error().Err(err).Msg("Failed to hijack Tempest client")
			return
		}
		mux := http.NewServeMux()
		mux.HandleFunc("/", httpHdl)

		log.Debug().Msgf("server listening at %v", b.lis.Addr())
		if err := http.Serve(b.lis, mux); err != nil {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	go func() {
		<-ctx.Done()
		log.Debug().Msg("Shutting down server...")
	}()

	return nil
}

package discord

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/rs/zerolog"
)

type Bot struct {
	logLevel zerolog.Level
	tmpst    *tempest.Client
	hdl      *handler.Handler
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

func New(hdl *handler.Handler, tmpst *tempest.Client) *Bot {
	return &Bot{
		tmpst: tmpst,
		hdl:   hdl,
	}
}

func (b *Bot) Initialize() error {
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

	guild, err := tempest.StringToSnowflake("948052547795574794")
	if err != nil {
		return err
	}
	if err := b.tmpst.SyncCommands([]tempest.Snowflake{guild}, nil, false); err != nil {
		return err
	}

	return nil
}

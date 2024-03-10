package discord

import (
	"context"
	"net"
	"net/http"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	comctx "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Bot struct {
	logLevel   zerolog.Level
	tmpst      *tempest.Client
	hdl        *handler.Handler
	mux        *http.ServeMux
	authorized []tempest.Snowflake
	guildIds   []tempest.Snowflake
}

func (b *Bot) Handler() *handler.Handler {
	return b.hdl
}

func (b *Bot) Tempest() *tempest.Client {
	return b.tmpst
}

func (b *Bot) LogLevel() zerolog.Level {
	return b.logLevel
}

func NewBot(hdl *handler.Handler, tmpst *tempest.Client, opts ...BotOption) *Bot {
	b := &Bot{
		logLevel: zerolog.InfoLevel,
		tmpst:    tmpst,
		hdl:      hdl,
	}
	for _, opt := range opts {
		opt(b)
	}

	return b
}

func (b *Bot) registerCommands() error {
	// TODO: This is gross and should be refactored
	if err := b.tmpst.RegisterCommand(command.NewServerHealthCommand(b.CommandPreflight(b.hdl.ServerHealthCheck))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewSelfHealthCheckCommand()); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeExecutionCommand(b.CommandPreflight(b.hdl.DescribeExecution))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewProvisionNodeCommand(b.CommandPreflight(b.hdl.ProvisionNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionProviderName:
			return b.hdl.ProviderNameAutoComplete(ctx, field, input.(string))
		case command.OptionTailnetName:
			return b.hdl.TailnetNameAutoComplete(ctx, field, input.(string))
		case command.OptionProviderLocation:
			return b.hdl.ProviderLocationAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDeprovisionNodeCommand(b.CommandPreflight(b.hdl.DeprovisionNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionNodeId:
			return b.hdl.NodeIdAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewListNodesCommand(b.CommandPreflight(b.hdl.ListNodes))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeNodeCommand(b.CommandPreflight(b.hdl.DescribeNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionNodeId:
			return b.hdl.NodeIdAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewStartNodeCommand(b.CommandPreflight(b.hdl.StartNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionNodeId:
			return b.hdl.NodeIdAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewStopNodeCommand(b.CommandPreflight(b.hdl.StopNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionNodeId:
			return b.hdl.NodeIdAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewListProvidersCommand(b.CommandPreflight(b.hdl.ListProviders))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeProviderCommand(b.CommandPreflight(b.hdl.DescribeProvider), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionProviderName:
			return b.hdl.ProviderNameAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewListTailnetsCommand(b.CommandPreflight(b.hdl.ListTailnets))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewListNotifiersCommand(b.CommandPreflight(b.hdl.ListNotifiers))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeTailnetCommand(b.CommandPreflight(b.hdl.DescribeTailnet), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case command.OptionTailnetName:
			return b.hdl.TailnetNameAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	return nil
}

func (b *Bot) Init() error {
	if err := b.registerCommands(); err != nil {
		return err
	}

	// set commands
	if err := b.tmpst.SyncCommands(nil, nil, false); err != nil {
		return err
	}
	// set guild commands
	if err := b.tmpst.SyncCommands(b.guildIds, nil, false); err != nil {
		return err
	}

	httpHdl, err := b.tmpst.Hijack()
	if err != nil {
		return err
	}
	b.mux = http.NewServeMux()
	b.mux.HandleFunc("/", httpHdl)

	return nil
}

func (b *Bot) HttpHandler() http.Handler {
	return b.mux
}

func (b *Bot) Serve(ctx context.Context, lis net.Listener) error {
	err := b.Init()
	if err != nil {
		return err
	}

	go func() {
		if err := http.Serve(lis, b.mux); err != nil {
			log.Error().Err(err).Msg("Server error")
		}
	}()

	go func() {
		<-ctx.Done()
		log.Debug().Msg("Shutting down server...")
	}()

	return nil
}

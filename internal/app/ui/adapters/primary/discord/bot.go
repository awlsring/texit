package discord

import (
	"context"
	"net"
	"net/http"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/command"
	comctx "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/context"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type Bot struct {
	logLevel   zerolog.Level
	tmpst      *tempest.Client
	hdl        *handler.Handler
	lis        net.Listener
	authorized []tempest.Snowflake
	guildId    *tempest.Snowflake
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

func NewBot(lis net.Listener, hdl *handler.Handler, tmpst *tempest.Client, lvl zerolog.Level, authorized []tempest.Snowflake, guild *tempest.Snowflake) *Bot {
	return &Bot{
		logLevel:   lvl,
		lis:        lis,
		tmpst:      tmpst,
		hdl:        hdl,
		authorized: authorized,
	}
}

func (b *Bot) registerCommands() error {
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
		case option.ProviderName:
			return b.hdl.ProviderNameAutoComplete(ctx, field, input.(string))
		case option.TailnetName:
			return b.hdl.TailnetNameAutoComplete(ctx, field, input.(string))
		case option.ProviderLocation:
			return b.hdl.ProviderLocationAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDeprovisionNodeCommand(b.CommandPreflight(b.hdl.DeprovisionNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case option.NodeId:
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
		case option.NodeId:
			return b.hdl.NodeIdAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewStartNodeCommand(b.CommandPreflight(b.hdl.StartNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case option.NodeId:
			return b.hdl.NodeIdAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewStopNodeCommand(b.CommandPreflight(b.hdl.StopNode), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case option.NodeId:
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
		case option.ProviderName:
			return b.hdl.ProviderNameAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewListTailnetsCommand(b.CommandPreflight(b.hdl.ListTailnets))); err != nil {
		return err
	}

	if err := b.tmpst.RegisterCommand(command.NewDescribeTailnetCommand(b.CommandPreflight(b.hdl.DescribeTailnet), b.AutoCompletePreflight(b.logLevel, func(ctx *comctx.CommandContext) []tempest.Choice {
		field, input := ctx.GetFocusedValue()
		switch field {
		case option.TailnetName:
			return b.hdl.TailnetNameAutoComplete(ctx, field, input.(string))
		}
		return nil
	}))); err != nil {
		return err
	}

	return nil
}

type CommandFunc func(*comctx.CommandContext)

func (b *Bot) auth(ctx *comctx.CommandContext) bool {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Checking authorization")

	if len(b.authorized) == 0 {
		log.Debug().Msg("No authorized users, allowing")
		return true
	}

	for _, id := range b.authorized {
		if ctx.Requester() == id {
			log.Debug().Msg("User is authorized")
			return true
		}
		if ctx.RequesterRoles() != nil {
			for _, r := range ctx.RequesterRoles() {
				if r == id {
					log.Debug().Msg("User is in authorized role")
					return true
				}
			}
		}
	}

	log.Debug().Msgf("User %d is not authorized", ctx.Requester())
	return false
}

func (b *Bot) CommandPreflight(comFunc CommandFunc) func(itx *tempest.CommandInteraction) {
	return func(itx *tempest.CommandInteraction) {
		ctx, err := comctx.InitContext(b.tmpst, itx, b.logLevel)
		if err != nil {
			return
		}
		if !b.auth(ctx) {
			_ = ctx.SendLinearReply("You are not authorized to use this command", true)
			return
		}
		log := ctx.Logger()
		log.Debug().Msg("Deferring command interaction")
		if err := itx.Defer(true); err != nil {
			log.Error().Err(err).Msg("Failed to defer command interaction")
			if err = ctx.SendLinearReply("Command failed with an unknown error!", true); err != nil {
				log.Error().Err(err).Msg("Failed to write bot response")
			}
			return
		}
		comFunc(ctx)
	}
}

type AutoCompleteFunc func(*comctx.CommandContext) []tempest.Choice

func (b *Bot) AutoCompletePreflight(logLevel zerolog.Level, comFunc AutoCompleteFunc) func(itx tempest.CommandInteraction) []tempest.Choice {
	return func(itx tempest.CommandInteraction) []tempest.Choice {
		ctx, err := comctx.InitContext(b.tmpst, &itx, logLevel)
		if err != nil {
			return nil
		}
		if !b.auth(ctx) {
			return nil
		}
		return comFunc(ctx)
	}
}

func (b *Bot) Start(ctx context.Context) error {
	log := logger.FromContext(ctx)

	if err := b.registerCommands(); err != nil {
		log.Error().Err(err).Msg("Failed to register commands")
		return err
	}

	var guilds []tempest.Snowflake
	if b.guildId != nil {
		guilds = append(guilds, *b.guildId)
	} else {
		guilds = nil
	}
	if err := b.tmpst.SyncCommands(guilds, nil, false); err != nil {
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

package command

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/rs/zerolog"
)

func NewServerHealthCommand(lvl zerolog.Level, tmpst *tempest.Client, hdl *handler.Handler) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "server-health",
		Description:         "Check the health of the texit server",
		SlashCommandHandler: CommandWrapper(lvl, tmpst, hdl.ServerHealthCheck),
	}
}

func NewSelfHealthCheckCommand(hdl *handler.Handler) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "self-health",
		Description:   "Check the health of the bot",
		SlashCommandHandler: func(itx *tempest.CommandInteraction) {
			itx.SendLinearReply("I'm healthy!", true)
		},
	}
}

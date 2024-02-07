package command

import (
	tempest "github.com/Amatsagu/Tempest"
)

func NewServerHealthCommand(slash func(itx *tempest.CommandInteraction)) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "server-health",
		Description:         "Check the health of the texit server",
		SlashCommandHandler: slash,
	}
}

func NewSelfHealthCheckCommand() tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "self-health",
		Description:   "Check the health of the bot",
		SlashCommandHandler: func(itx *tempest.CommandInteraction) {
			_ = itx.SendLinearReply("I'm healthy!", true)
		},
	}
}

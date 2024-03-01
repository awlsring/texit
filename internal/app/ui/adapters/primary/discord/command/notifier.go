package command

import (
	tempest "github.com/Amatsagu/Tempest"
)

func NewListNotifiersCommand(slash func(itx *tempest.CommandInteraction)) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "list-notifiers",
		Description:         "List all notifiers.",
		SlashCommandHandler: slash,
	}
}

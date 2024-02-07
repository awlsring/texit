package command

import (
	tempest "github.com/Amatsagu/Tempest"
)

func NewListTailnetsCommand(slash func(itx *tempest.CommandInteraction)) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "list-tailnets",
		Description:         "List all tailnets.",
		SlashCommandHandler: slash,
	}
}

func NewDescribeTailnetCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "describe-tailnet",
		Description:   "Describes a tailnet.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         OptionTailnetName,
				Description:  "The name of the tailnet to describe",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

package command

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
)

func NewListProvidersCommand(slash func(itx *tempest.CommandInteraction)) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "list-providers",
		Description:         "List all providers.",
		SlashCommandHandler: slash,
	}
}

func NewDescribeProviderCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "describe-provider",
		Description:   "Describes a provider.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.ProviderName,
				Description:  "The name of the provider to describe",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

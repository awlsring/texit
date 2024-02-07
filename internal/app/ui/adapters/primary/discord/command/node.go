package command

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
)

func NewProvisionNodeCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "create-exit-node",
		Description:   "Create an Exit Node on a tailnet and provider.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.ProviderName,
				Description:  "The name of the provider to create the exit node in",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.TailnetName,
				Description:  "The tailnet to add the exit node to.",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.ProviderLocation,
				Description:  "The location of the provider to create the exit node on",
				Required:     true,
				MinValue:     3,
				AutoComplete: true,
			},
			{
				Type:        tempest.BOOLEAN_OPTION_TYPE,
				Name:        option.Ephemeral,
				Description: "Whether the created exit node should be ephemeral or not. Defaults to false.",
				Required:    false,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

func NewDeprovisionNodeCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "delete-exit-node",
		Description:   "Deletes an Exit Node.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.NodeId,
				Description:  "The ID of the node to delete",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

func NewListNodesCommand(slash func(itx *tempest.CommandInteraction)) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "list-exit-nodes",
		Description:         "List all Exit Nodes.",
		SlashCommandHandler: slash,
	}
}

func NewDescribeNodeCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "describe-exit-node",
		Description:   "Describes an Exit Node.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.NodeId,
				Description:  "The ID of the node to describe",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

func NewStartNodeCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "start-exit-node",
		Description:   "Starts an Exit Node.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.NodeId,
				Description:  "The ID of the node to start",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

func NewStopNodeCommand(slash func(itx *tempest.CommandInteraction), auto func(itx tempest.CommandInteraction) []tempest.Choice) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "stop-exit-node",
		Description:   "Stops an Exit Node.",
		Options: []tempest.CommandOption{
			{
				Type:         tempest.STRING_OPTION_TYPE,
				Name:         option.NodeId,
				Description:  "The ID of the node to stop",
				Required:     true,
				MinValue:     1,
				AutoComplete: true,
			},
		},
		SlashCommandHandler: slash,
		AutoCompleteHandler: auto,
	}
}

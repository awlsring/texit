package command

import (
	"context"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
	"github.com/rs/zerolog"
)

func NewProvisionNodeCommand(lvl zerolog.Level, tmpst *tempest.Client, hdl *handler.Handler) tempest.Command {
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
				Type:        tempest.STRING_OPTION_TYPE,
				Name:        option.ProviderLocation,
				Description: "The location of the provider to create the exit node on",
				Required:    true,
				MinValue:    3,
			},
			{
				Type:        tempest.BOOLEAN_OPTION_TYPE,
				Name:        option.Ephemeral,
				Description: "Whether the created exit node should be executed or not. Defaults to false.",
				Required:    false,
			},
		},
		SlashCommandHandler: CommandWrapper(lvl, tmpst, hdl.ProvisionNode),
		AutoCompleteHandler: AutoCompleteWrapper(lvl, func(ctx context.Context, itx tempest.CommandInteraction) []tempest.Choice {
			field, input := itx.GetFocusedValue()
			switch field {
			case option.ProviderName:
				return hdl.ProviderNameAutoComplete(ctx, itx, field, input.(string))
			case option.TailnetName:
				return hdl.TailnetNameAutoComplete(ctx, itx, field, input.(string))
			}
			return nil
		}),
	}
}

func NewDeprovisionNodeCommand(lvl zerolog.Level, tmpst *tempest.Client, hdl *handler.Handler) tempest.Command {
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
		SlashCommandHandler: CommandWrapper(lvl, tmpst, hdl.DeprovisionNode),
		AutoCompleteHandler: AutoCompleteWrapper(lvl, func(ctx context.Context, itx tempest.CommandInteraction) []tempest.Choice {
			field, input := itx.GetFocusedValue()
			switch field {
			case option.NodeId:
				return hdl.NodeIdAutoComplete(ctx, itx, field, input.(string))
			}
			return nil
		}),
	}
}

func NewListNodesCommand(lvl zerolog.Level, tmpst *tempest.Client, hdl *handler.Handler) tempest.Command {
	return tempest.Command{
		AvailableInDM:       true,
		Name:                "list-exit-nodes",
		Description:         "List all Exit Nodes.",
		SlashCommandHandler: CommandWrapper(lvl, tmpst, hdl.ListNodes),
	}
}

func NewDescribeNodeCommand(lvl zerolog.Level, tmpst *tempest.Client, hdl *handler.Handler) tempest.Command {
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
		SlashCommandHandler: CommandWrapper(lvl, tmpst, hdl.DescribeNode),
		AutoCompleteHandler: AutoCompleteWrapper(lvl, func(ctx context.Context, itx tempest.CommandInteraction) []tempest.Choice {
			field, input := itx.GetFocusedValue()
			switch field {
			case option.NodeId:
				return hdl.NodeIdAutoComplete(ctx, itx, field, input.(string))
			}
			return nil
		}),
	}
}

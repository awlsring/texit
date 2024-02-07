package command

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
)

func NewDescribeExecutionCommand(slash func(itx *tempest.CommandInteraction)) tempest.Command {
	return tempest.Command{
		AvailableInDM: true,
		Name:          "describe-execution",
		Description:   "Describe an execution by its ID",
		Options: []tempest.CommandOption{
			{
				Type:        tempest.STRING_OPTION_TYPE,
				Name:        option.ExecutionId,
				Description: "The ID of the execution to describe",
				Required:    true,
				MinValue:    32,
			},
		},
		SlashCommandHandler: slash,
	}
}

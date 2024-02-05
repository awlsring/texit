package command

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/handler"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/option"
	"github.com/rs/zerolog"
)

func NewDescribeExecutionCommand(lvl zerolog.Level, tmpst *tempest.Client, hdl *handler.Handler) tempest.Command {
	return tempest.Command{
		Name:        "describe-execution",
		Description: "Describe an execution by its ID",
		Options: []tempest.CommandOption{
			{
				Type:        tempest.STRING_OPTION_TYPE,
				Name:        option.ExecutionId,
				Description: "The ID of the execution to describe",
				Required:    true,
				MinValue:    32,
			},
		},
		SlashCommandHandler: CommandWrapper(lvl, tmpst, hdl.DescribeExecution),
	}
}

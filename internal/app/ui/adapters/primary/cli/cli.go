package cli

import (
	"log"
	"os"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/adapters/primary/cli/handler"
	"github.com/urfave/cli/v2"
)

type CLI struct {
	app *cli.App
}

func New(hdl *handler.Handler) *CLI {
	app := &cli.App{
		// Flags: []cli.Flag{
		// 	&cli.StringFlag{
		// 		Name:    "config",
		// 		Aliases: []string{"c"},
		// 		Usage:   "Load configuration from `FILE`",
		// 	},
		// },
		Commands: []*cli.Command{
			{
				Name:   "health",
				Usage:  "Check the health of the server",
				Action: hdl.HealthCheck,
			},
		},
	}

	return &CLI{
		app: app,
	}
}

func (c *CLI) Run() error {
	if err := c.app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	return nil
}

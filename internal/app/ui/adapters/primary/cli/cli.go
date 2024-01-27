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
		Commands: []*cli.Command{
			{
				Name:   "health",
				Usage:  "Check the health of the server",
				Action: hdl.HealthCheck,
			},
			{
				Name: "provider",
				Subcommands: []*cli.Command{
					{
						Name:   "default",
						Usage:  "Gets the default provider",
						Action: hdl.GetDefaultProvider,
					},
					{
						Name:   "describe",
						Usage:  "Describes a provider",
						Action: hdl.GetProvider,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     handler.ProviderNameFlag,
								Aliases:  []string{"n"},
								Usage:    "The name of the provider to describe",
								Required: true,
							},
						},
					},
					{
						Name:   "list",
						Usage:  "Lists all providers",
						Action: hdl.ListProviders,
					},
				},
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

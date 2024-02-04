package cli

import (
	"log"
	"os"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/handler"
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
				Name: "node",
				Subcommands: []*cli.Command{
					{
						Name:        "list",
						Description: "Lists all exit nodes",
						Action:      hdl.ListNodes,
					},
					{
						Name:        "describe",
						Description: "Describes an exit node",
						Action:      hdl.DescribeNode,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.NodeId,
								Aliases:  []string{"i"},
								Usage:    "The id of the node to describe",
								Required: true,
							},
						},
					},
					{
						Name:        "status",
						Description: "Describes the status of an exit node",
						Action:      hdl.GetNodeStatus,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.NodeId,
								Aliases:  []string{"i"},
								Usage:    "The id of the node to get the status of",
								Required: true,
							},
						},
					},
					{
						Name:        "provision",
						Description: "Provision a new exit node on given provider in a given location",
						Action:      hdl.ProvisionNode,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.Provider,
								Aliases:  []string{"p"},
								Usage:    "The name of the provider to provision the node on",
								Required: true,
							},
							&cli.StringFlag{
								Name:     flag.Tailnet,
								Aliases:  []string{"t"},
								Usage:    "The tailnet to add the node to",
								Required: true,
							},
							&cli.StringFlag{
								Name:     flag.ProviderLocation,
								Aliases:  []string{"l"},
								Usage:    "The location to provision the node in",
								Required: true,
							},
							&cli.BoolFlag{
								Name:    flag.EphemeralNode,
								Aliases: []string{"e"},
								Usage:   "Whether or not the node should be ephemeral",
							},
						},
					},
					{
						Name:        "deprovision",
						Description: "Deprovision an exit node",
						Action:      hdl.DeprovisionNode,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.NodeId,
								Aliases:  []string{"n"},
								Usage:    "The id of the node to deprovision",
								Required: true,
							},
						},
					},
					{
						Name:        "start",
						Description: "Start an exit node",
						Action:      hdl.StartNode,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.NodeId,
								Aliases:  []string{"n"},
								Usage:    "The id of the node to start",
								Required: true,
							},
						},
					},
					{
						Name:        "stop",
						Description: "Stop an exit node",
						Action:      hdl.StopNode,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.NodeId,
								Aliases:  []string{"n"},
								Usage:    "The id of the node to stop",
								Required: true,
							},
						},
					},
				},
			},
			{
				Name: "execution",
				Subcommands: []*cli.Command{
					{
						Name:        "describe",
						Description: "Describes an execution",
						Action:      hdl.DescribeExecution,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.ExecutionId,
								Aliases:  []string{"i"},
								Usage:    "The id of the execution to describe",
								Required: true,
							},
						},
					},
				},
			},
			{
				Name: "provider",
				Subcommands: []*cli.Command{
					{
						Name:   "describe",
						Usage:  "Describes a provider",
						Action: hdl.GetProvider,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.ProviderName,
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
					{
						Name:   "init",
						Usage:  "Help command to initialize a provider for use with Texit",
						Action: hdl.ProviderInit,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.ProviderType,
								Aliases:  []string{"t"},
								Usage:    "The type of provider to initialize",
								Required: true,
							},
						},
					},
				},
			},
			{
				Name: "tailnet",
				Subcommands: []*cli.Command{
					{
						Name:   "describe",
						Usage:  "Describes a tailnet",
						Action: hdl.GetTailnet,
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     flag.TailnetName,
								Aliases:  []string{"n"},
								Usage:    "The name of the tailnet to describe",
								Required: true,
							},
						},
					},
					{
						Name:   "list",
						Usage:  "Lists all tailnets",
						Action: hdl.ListTailnets,
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

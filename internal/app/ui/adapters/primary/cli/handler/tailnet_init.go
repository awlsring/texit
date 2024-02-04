package handler

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/urfave/cli/v2"
)

func (h *Handler) TailnetInit(c *cli.Context) error {
	t, err := tailnet.TypeFromString(c.String(flag.TailnetType))
	if err != nil {
		return err
	}

	switch t {
	case tailnet.TypeHeadscale:
		return h.initHeadscaleTailnet(c)
	case tailnet.TypeTailscale:
		return h.initTailscaleTailnet(c)
	default:
		return fmt.Errorf("unknown tailnet type: %s", t.String())
	}
}

func (h *Handler) initTailscaleTailnet(c *cli.Context) error {
	fmt.Println("To use a Tailscale tailnet, you must have a tailnet created on Tailscale and access to the Admin panel.")
	fmt.Println("")
	fmt.Println("Create an access key by going to https://login.tailscale.com/admin/settings/keys and clicking 'Generate Access Token'.")
	fmt.Println("")
	fmt.Println("Once you have the access key, add the following to your texit config.yaml")
	fmt.Print(`
tailnets:
  ...
  - apiKey: <THE KEY YOU JUST MADE>
    tailnet: <THE NAME OF YOUR TAILNET>
    type: "tailscale"
    user: <YOUR TAILSCALE USER>
`)
	return nil
}

func (h *Handler) initHeadscaleTailnet(c *cli.Context) error {
	fmt.Println("To use a Headscale tailnet, you must have a Headscale server running and the Headscale CLI configured.")
	fmt.Println("")
	fmt.Println("Create an API key with the following command:")
	fmt.Println("")
	fmt.Println("headscale apikey create")
	fmt.Println("")
	fmt.Println("Once you have the api key, add the following to your texit config.yaml")
	fmt.Print(`
tailnets:
  ...
  - apiKey: <THE KEY YOU JUST MADE>
    tailnet: <THE NAME OF YOUR TAILNET>
    type: "headscale"
    user: <YOUR HEADSCALE USER>
`)
	return nil
}

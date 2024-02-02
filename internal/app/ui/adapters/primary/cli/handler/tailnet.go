package handler

import (
	"context"
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/urfave/cli/v2"
)

func generateTailnetString(tn *tailnet.Tailnet) string {
	return fmt.Sprintf("Tailnet - Name: %s | Type: %s | ControlServer: %s", tn.Name, tn.Type.String(), tn.ControlServer.String())
}

func (h *Handler) GetTailnet(c *cli.Context) error {
	name, err := tailnet.IdentifierFromString(c.String(flag.ProviderName))
	if err != nil {
		return err
	}
	tn, err := h.apiSvc.GetTailnet(context.Background(), name)
	if err != nil {
		return err
	}
	fmt.Println(generateTailnetString(tn))
	return nil
}

func (h *Handler) ListTailnets(ctx *cli.Context) error {
	tns, err := h.apiSvc.ListTailnets(context.Background())
	if err != nil {
		return err
	}

	if len(tns) == 0 {
		fmt.Println("No tailnets found")
		return nil
	}

	fmt.Printf("Tailnets: %d\n", len(tns))
	fmt.Println("==========================")
	for _, tn := range tns {
		fmt.Println(generateTailnetString(tn))
	}
	return nil
}

package handler

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
)

func (h *Handler) ListNotifiers(ctx *cli.Context) error {
	nts, err := h.apiSvc.ListNotifiers(context.Background())
	if err != nil {
		return err
	}

	if len(nts) == 0 {
		fmt.Println("No notifiers found")
		return nil
	}

	fmt.Printf("Notifiers: %d\n", len(nts))
	fmt.Println("==========================")
	for _, n := range nts {
		fmt.Printf("- Name: %s | Type: %s | Endpoint: %s\n", n.Name.String(), n.Type.String(), n.Endpoint.String())
	}
	return nil
}

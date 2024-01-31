package handler

import (
	"context"
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/urfave/cli/v2"
)

func (h *Handler) GetProvider(c *cli.Context) error {
	name, err := provider.IdentifierFromString(c.String(flag.ProviderName))
	if err != nil {
		return err
	}
	prov, err := h.apiSvc.GetProvider(context.Background(), name)
	if err != nil {
		return err
	}
	fmt.Printf("Provider - Name: %s | Platform: %s", prov.Name, prov.Platform.String())
	return nil
}

func (h *Handler) ListProviders(ctx *cli.Context) error {
	provs, err := h.apiSvc.ListProviders(context.Background())
	if err != nil {
		return err
	}

	if len(provs) == 0 {
		fmt.Println("No providers found")
		return nil
	}

	fmt.Printf("Providers: %d\n", len(provs))
	fmt.Println("==========================")
	for _, prov := range provs {
		fmt.Printf("Provider - Name: %s | Platform: %s", prov.Name, prov.Platform.String())
	}
	return nil
}

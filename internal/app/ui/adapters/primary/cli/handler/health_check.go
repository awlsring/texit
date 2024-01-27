package handler

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"
)

func (h *Handler) HealthCheck(ctx *cli.Context) error {
	err := h.apiSvc.CheckServerHealth(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("Server is healthy!")
	return nil
}

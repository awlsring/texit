package handler

import (
	"context"

	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) CreateNode(context.Context, *teen.CreateNodeRequest) (*teen.CreateNodeResponse, error) {
	panic("implement me")
}

package handler

import "github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/ports/service"

type Handler struct {
	apiSvc service.Api
}

func New(svc service.Api) *Handler {
	return &Handler{
		apiSvc: svc,
	}
}

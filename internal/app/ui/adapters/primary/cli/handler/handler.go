package handler

import "github.com/awlsring/texit/internal/app/ui/ports/service"

type Handler struct {
	apiSvc service.Api
}

func New(svc service.Api) *Handler {
	return &Handler{
		apiSvc: svc,
	}
}

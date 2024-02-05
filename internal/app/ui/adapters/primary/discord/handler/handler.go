package handler

import (
	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
)

type Handler struct {
	apiSvc service.Api
}

func New(svc service.Api) *Handler {
	return &Handler{
		apiSvc: svc,
	}
}

func (h *Handler) editResponseMessage(itx *tempest.CommandInteraction, message string) error {
	return itx.EditReply(tempest.ResponseMessageData{
		Content: message,
	}, true)
}

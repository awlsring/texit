package notification

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/service"
)

type Service struct {
	gateways []gateway.Notification
}

func NewService(gateways []gateway.Notification) service.Notification {
	return &Service{
		gateways: gateways,
	}
}

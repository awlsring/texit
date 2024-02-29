package notification

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
)

type Service struct {
	notifiers []*notification.Notifier
	gateways  []gateway.Notification
}

func NewService(gateways map[string]gateway.Notification) (service.Notification, error) {
	gws := []gateway.Notification{}
	nts := []*notification.Notifier{}
	for n, g := range gateways {
		gws = append(gws, g)
		name, err := notification.IdentifierFromString(n)
		if err != nil {
			return nil, err
		}
		nts = append(nts, notification.NewNotifier(name, g.Type(), g.Endpoint()))
	}

	return &Service{
		notifiers: nts,
		gateways:  gws,
	}, nil
}

package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) ListNotifiers(ctx context.Context) ([]*notification.Notifier, error) {
	resp, err := g.client.ListNotifiers(ctx)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	nts := []*notification.Notifier{}
	for _, n := range resp.Summaries {
		node, err := SummaryToNotifiers(n)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		nts = append(nts, node)
	}

	return nts, nil
}

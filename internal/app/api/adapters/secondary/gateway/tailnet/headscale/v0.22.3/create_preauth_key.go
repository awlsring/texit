package headscale_v0_22_3_gateway

import (
	"context"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/client/headscale_service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/models"
	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
)

func (g *HeadscaleGateway) CreatePreauthKey(ctx context.Context, ephemeral bool) (tailnet.PreauthKey, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("creating headscale preauth key")

	log.Debug().Msg("creating headscale preauth key request")
	request := headscale_service.NewHeadscaleServiceCreatePreAuthKeyParams()
	request.SetContext(ctx)
	body := &models.V1CreatePreAuthKeyRequest{
		User:       g.user,
		Reusable:   false,
		Ephemeral:  ephemeral,
		Expiration: strfmt.DateTime(time.Now().Add(oneYearExpiration)),
	}
	request.SetBody(body)

	log.Debug().Msg("calling headscale")
	resp, err := g.client.HeadscaleServiceCreatePreAuthKey(request)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("headscale preauth key created")
	err = resp.Payload.Validate(strfmt.Default)
	if err != nil {
		return "", err
	}

	log.Debug().Msg("parsing headscale preauth key")
	key, err := tailnet.PreauthKeyFromString(resp.Payload.PreAuthKey.Key)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse headscale preauth key")
	}

	log.Debug().Msg("headscale preauth key parsed")
	return key, nil
}

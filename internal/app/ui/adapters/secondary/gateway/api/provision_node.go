package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func setNodeSize(size node.Size) texit.OptNodeSize {
	var s texit.NodeSize
	switch size {
	case node.SizeSmall:
		s = texit.NodeSizeSmall
	case node.SizeMedium:
		s = texit.NodeSizeMedium
	case node.SizeLarge:
		s = texit.NodeSizeLarge
	}
	return texit.NewOptNodeSize(s)
}

func (g *ApiGateway) ProvisionNode(ctx context.Context, prov provider.Identifier, loc provider.Location, tn tailnet.Identifier, sz node.Size, eph bool) (workflow.ExecutionIdentifier, error) {
	req := &texit.ProvisionNodeRequestContent{
		Provider:  prov.String(),
		Location:  loc.String(),
		Tailnet:   tn.String(),
		Ephemeral: texit.NewOptBool(eph),
		Size:      setNodeSize(sz),
	}

	resp, err := g.client.ProvisionNode(ctx, req)
	if err != nil {
		return "", errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	switch resp := resp.(type) {
	case *texit.ProvisionNodeResponseContent:
		id, err := workflow.ExecutionIdentifierFromString(resp.Execution)
		if err != nil {
			return "", errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}

		return id, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return "", errors.Wrap(gateway.ErrResourceNotFoundError, resp.Message)
	case *texit.InvalidInputErrorResponseContent:
		return "", errors.Wrap(gateway.ErrInvalidInputError, resp.Message)
	default:
		return "", gateway.ErrInternalServerError
	}
}

package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/ports/gateway"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"google.golang.org/grpc/metadata"
)

type ApiGateway struct {
	apiKey string
	client v1.TailscaleEphemeralExitNodesServiceClient
}

func New(apiKey string, client v1.TailscaleEphemeralExitNodesServiceClient) gateway.Api {
	return &ApiGateway{
		apiKey: apiKey,
		client: client,
	}
}

func (s *ApiGateway) setAuthInContext(ctx context.Context) context.Context {
	md := metadata.Pairs("api-key", s.apiKey)
	return metadata.NewOutgoingContext(ctx, md)
}

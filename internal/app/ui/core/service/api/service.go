package api

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/ports/service"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"google.golang.org/grpc/metadata"
)

type Service struct {
	apiKey string
	client v1.TailscaleEphemeralExitNodesServiceClient
}

func NewService(apiKey string, client v1.TailscaleEphemeralExitNodesServiceClient) service.Api {
	return &Service{
		apiKey: apiKey,
		client: client,
	}
}

func (s *Service) setAuthInContext(ctx context.Context) context.Context {
	md := metadata.Pairs("api-key", s.apiKey)
	return metadata.NewOutgoingContext(ctx, md)
}

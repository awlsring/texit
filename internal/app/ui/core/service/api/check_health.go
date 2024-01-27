package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Service) CheckServerHealth(ctx context.Context) error {
	ctx = s.setAuthInContext(ctx)
	_, err := s.client.HealthCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return err
	}

	return nil
}

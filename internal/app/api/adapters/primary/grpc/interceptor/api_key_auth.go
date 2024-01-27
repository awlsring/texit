package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func NewStaticAuthKeyInterceptor(key string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}
		fmt.Println(md)

		apiKeys, ok := md["api-key"]
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "api key is not provided")
		}

		if len(apiKeys) != 1 {
			return nil, status.Errorf(codes.Unauthenticated, "expected 1 api key, got %d", len(apiKeys))
		}

		if apiKeys[0] != key {
			return nil, status.Errorf(codes.Unauthenticated, "invalid api key")
		}

		return handler(ctx, req)
	}
}

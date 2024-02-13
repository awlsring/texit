package clients

import (
	"context"

	"github.com/awlsring/texit/pkg/gen/texit"
)

type TexitSecurity struct {
	key string
}

func (s TexitSecurity) SmithyAPIHttpApiKeyAuth(ctx context.Context, operationName string) (texit.SmithyAPIHttpApiKeyAuth, error) {
	return texit.SmithyAPIHttpApiKeyAuth{
		APIKey: s.key,
	}, nil
}

func CreateTexitClient(address string, key string) (texit.Invoker, error) {
	return texit.NewClient(address, TexitSecurity{key: key})
}

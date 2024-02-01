package auth

import (
	"context"
	"errors"

	"github.com/awlsring/texit/pkg/gen/texit"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

type StaticKeySecurityHandlerOption func(*StaticKeySecurityHandler)

type StaticKeySecurityHandler struct {
	keys             []string
	noAuthOperations []string
}

func WithNoAuthOperations(operations []string) StaticKeySecurityHandlerOption {
	return func(h *StaticKeySecurityHandler) {
		h.noAuthOperations = operations
	}
}

func NewSecurityHandler(keys []string, opts ...StaticKeySecurityHandlerOption) texit.SecurityHandler {
	sec := &StaticKeySecurityHandler{
		keys: keys,
	}
	for _, opt := range opts {
		opt(sec)
	}
	return sec
}

func stringInList(s string, list []string) bool {
	for _, item := range list {
		if s == item {
			return true
		}
	}
	return false
}

func (h *StaticKeySecurityHandler) isNoAuthOperation(operationName string) bool {
	return stringInList(operationName, h.noAuthOperations)
}

func (h *StaticKeySecurityHandler) isValidKey(key string) bool {
	return stringInList(key, h.keys)
}

func (h *StaticKeySecurityHandler) HandleSmithyAPIHttpApiKeyAuth(ctx context.Context, operationName string, t texit.SmithyAPIHttpApiKeyAuth) (context.Context, error) {
	key := t.GetAPIKey()

	if h.isNoAuthOperation(operationName) {
		return ctx, nil
	}

	if h.isValidKey(key) {
		return ctx, nil
	}

	return ctx, ErrInvalidKey
}

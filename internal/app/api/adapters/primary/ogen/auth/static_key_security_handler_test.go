package auth

import (
	"context"
	"testing"

	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestStaticKeySecurityHandler(t *testing.T) {
	keys := []string{"key1", "key2"}
	noAuthOperations := []string{"operation1", "operation2"}

	handler := NewSecurityHandler(keys, WithNoAuthOperations(noAuthOperations)).(*StaticKeySecurityHandler)

	t.Run("stringInList", func(t *testing.T) {
		assert.True(t, stringInList("key1", keys))
		assert.False(t, stringInList("key3", keys))
	})

	t.Run("isNoAuthOperation", func(t *testing.T) {
		assert.True(t, handler.isNoAuthOperation("operation1"))
		assert.False(t, handler.isNoAuthOperation("operation3"))
	})

	t.Run("isValidKey", func(t *testing.T) {
		assert.True(t, handler.isValidKey("key1"))
		assert.False(t, handler.isValidKey("key3"))
	})

	t.Run("HandleSmithyAPIHttpApiKeyAuth", func(t *testing.T) {
		ctx := context.Background()
		auth := texit.SmithyAPIHttpApiKeyAuth{APIKey: "key1"}

		newCtx, err := handler.HandleSmithyAPIHttpApiKeyAuth(ctx, "operation1", auth)
		assert.NoError(t, err)
		assert.Equal(t, ctx, newCtx)

		newCtx, err = handler.HandleSmithyAPIHttpApiKeyAuth(ctx, "operation3", auth)
		assert.NoError(t, err)
		assert.Equal(t, ctx, newCtx)

		auth = texit.SmithyAPIHttpApiKeyAuth{APIKey: "key3"}
		newCtx, err = handler.HandleSmithyAPIHttpApiKeyAuth(ctx, "operation3", auth)
		assert.Error(t, err)
		assert.Equal(t, ErrInvalidKey, err)
		assert.Equal(t, ctx, newCtx)
	})
}

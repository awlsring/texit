package platform_aws_ecs

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/patrickmn/go-cache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

type MockOption struct{}

func (m *MockClient) SomeMethod() {
	m.Called()
}

func TestGetClientForLocation(t *testing.T) {
	ctx := context.Background()
	ch := cache.New(5*time.Minute, 10*time.Minute)
	loc := provider.Location("test-location")
	creds := credentials.NewStaticCredentialsProvider("test-id", "test-secret", "test-token")

	clientFunc := func(cfg aws.Config, opts ...func(*MockOption)) *MockClient {
		return &MockClient{}
	}

	client, err := getClientForLocation(ctx, clientFunc, ch, loc, creds)

	assert.NoError(t, err)
	assert.IsType(t, &MockClient{}, client)

	client2, err := getClientForLocation(ctx, clientFunc, ch, loc, creds)

	assert.NoError(t, err)
	assert.Equal(t, client, client2)
}

package tailscale_gateway

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

const (
	testUser = "test-user"
)

func TestUpdateAcl(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	mockAcl := &tailscale.ACL{
		AutoApprovers: &tailscale.ACLAutoApprovers{
			ExitNode: []string{},
		},
		TagOwners: map[string][]string{},
	}

	mockClient.EXPECT().ACL(ctx).Return(mockAcl, nil)
	mockClient.EXPECT().ValidateACL(ctx, *mockAcl).Return(nil)
	mockClient.EXPECT().SetACL(ctx, *mockAcl).Return(nil)

	g := &TailscaleGateway{
		client: mockClient,
		user:   testUser,
	}

	err := g.updateAcl(ctx)

	assert.NoError(t, err)
	assert.Contains(t, mockAcl.AutoApprovers.ExitNode, tagCloudExitNode)
	assert.Contains(t, mockAcl.TagOwners[tagCloudExitNode], g.user)
}

func TestUpdateAcl_NoUpdateNeeded(t *testing.T) {
	ctx := context.Background()

	mockClient := mocks.NewMockTailscale_interfaces(t)
	mockAcl := &tailscale.ACL{
		AutoApprovers: &tailscale.ACLAutoApprovers{
			ExitNode: []string{tagCloudExitNode},
		},
		TagOwners: map[string][]string{tagCloudExitNode: {testUser}},
	}
	mockClient.EXPECT().ACL(ctx).Return(mockAcl, nil)

	g := &TailscaleGateway{
		client: mockClient,
		user:   testUser,
	}

	err := g.updateAcl(ctx)

	assert.NoError(t, err)
	assert.Contains(t, mockAcl.AutoApprovers.ExitNode, tagCloudExitNode)
	assert.Contains(t, mockAcl.TagOwners[tagCloudExitNode], g.user)
}

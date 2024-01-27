package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

type AwsAccountClient interface {
	GetRegionOptStatus(ctx context.Context, params *account.GetRegionOptStatusInput, optFns ...func(*account.Options)) (*account.GetRegionOptStatusOutput, error)
	EnableRegion(ctx context.Context, params *account.EnableRegionInput, optFns ...func(*account.Options)) (*account.EnableRegionOutput, error)
}

type EcsClient interface {
	CreateCluster(ctx context.Context, params *ecs.CreateClusterInput, optFns ...func(*ecs.Options)) (*ecs.CreateClusterOutput, error)
	DescribeClusters(ctx context.Context, params *ecs.DescribeClustersInput, optFns ...func(*ecs.Options)) (*ecs.DescribeClustersOutput, error)
	ListClusters(ctx context.Context, params *ecs.ListClustersInput, optFns ...func(*ecs.Options)) (*ecs.ListClustersOutput, error)
	CreateTaskSet(ctx context.Context, params *ecs.CreateTaskSetInput, optFns ...func(*ecs.Options)) (*ecs.CreateTaskSetOutput, error)
	CreateService(ctx context.Context, params *ecs.CreateServiceInput, optFns ...func(*ecs.Options)) (*ecs.CreateServiceOutput, error)
	DescribeServices(ctx context.Context, params *ecs.DescribeServicesInput, optFns ...func(*ecs.Options)) (*ecs.DescribeServicesOutput, error)
	RegisterTaskDefinition(ctx context.Context, params *ecs.RegisterTaskDefinitionInput, optFns ...func(*ecs.Options)) (*ecs.RegisterTaskDefinitionOutput, error)
	RunTask(ctx context.Context, params *ecs.RunTaskInput, optFns ...func(*ecs.Options)) (*ecs.RunTaskOutput, error)
	DeleteTaskDefinitions(ctx context.Context, params *ecs.DeleteTaskDefinitionsInput, optFns ...func(*ecs.Options)) (*ecs.DeleteTaskDefinitionsOutput, error)
	StopTask(ctx context.Context, params *ecs.StopTaskInput, optFns ...func(*ecs.Options)) (*ecs.StopTaskOutput, error)
}

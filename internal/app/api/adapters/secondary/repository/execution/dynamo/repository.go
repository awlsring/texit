package dynamo_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoExecutionRepository struct {
	table  string
	client *dynamodb.Client
}

func (*DynamoExecutionRepository) Init(ctx context.Context) error {
	return nil
}

func (r *DynamoExecutionRepository) Close() {
	return
}

func New(table string, client *dynamodb.Client) repository.Execution {
	return &DynamoExecutionRepository{
		table:  table,
		client: client,
	}
}

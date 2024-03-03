package dynamo_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoNodeRepository struct {
	table  string
	client *dynamodb.Client
}

func (*DynamoNodeRepository) Init(ctx context.Context) error {
	return nil
}

func (r *DynamoNodeRepository) Close() {}

func New(table string, client *dynamodb.Client) repository.Node {
	return &DynamoNodeRepository{
		table:  table,
		client: client,
	}
}

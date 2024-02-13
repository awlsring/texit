package dynamo_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (d *DynamoExecutionRepository) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting execution from DynamoDB")

	key := map[string]types.AttributeValue{
		AttributeIdentifier: &types.AttributeValueMemberS{Value: id.String()},
	}

	log.Debug().Interface("key", key).Msg("Getting execution from DynamoDB")
	result, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &d.table,
		Key:       key,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error getting execution from DynamoDB")
		return nil, err
	}
	if len(result.Item) == 0 {
		log.Warn().Msg("Node not found")
		return nil, repository.ErrExecutionNotFound
	}

	var entry ExecutionDdbRecord
	log.Debug().Interface("result", result).Msg("Unmarshalling execution from DynamoDB")
	err = attributevalue.UnmarshalMap(result.Item, &entry)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling execution from DynamoDB")
		return nil, err
	}

	log.Debug().Interface("entry", entry).Msg("Unmarshalled execution from DynamoDB")
	return entry.ToExecution(), nil
}

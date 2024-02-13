package dynamo_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (d *DynamoExecutionRepository) CreateExecution(ctx context.Context, ex *workflow.Execution) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating execution in DynamoDB")

	record := recordFromExecution(ex)
	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling execution to DynamoDB")
		return err
	}
	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &d.table, Item: item,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error putting execution to Dynamo")
		return err
	}

	log.Debug().Msg("Created execution in DynamoDB")
	return nil
}

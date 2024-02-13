package dynamo_execution_repository

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	TtlPeriod = 24 * time.Hour
)

func (d *DynamoExecutionRepository) CloseExecution(ctx context.Context, id workflow.ExecutionIdentifier, result workflow.Status, output workflow.SerializedExecutionResult) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Updating execution in DynamoDB")

	key := map[string]types.AttributeValue{
		AttributeIdentifier: &types.AttributeValueMemberS{Value: id.String()},
	}

	now := time.Now()

	log.Debug().Msg("Building update expression")
	update := expression.Set(expression.Name(AttributeStatus), expression.Value(result.String()))
	update.Set(expression.Name(AttributeUpdatedAt), expression.Value(now))
	update.Set(expression.Name(AttributeFinishedAt), expression.Value(now))
	update.Set(expression.Name(AttributeResults), expression.Value(output.String()))
	update.Set(expression.Name(AttributeTtl), expression.Value(time.Now().Add(TtlPeriod).Unix()))
	expr, err := expression.NewBuilder().WithUpdate(update).Build()
	if err != nil {
		log.Error().Err(err).Msg("Error building update expression")
		return err
	}

	log.Debug().Msg("Updating execution in DynamoDB")
	_, err = d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 &d.table,
		Key:                       key,
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Error updating execution in Dynamo")
		return err
	}

	log.Debug().Msg("Updated execution in DynamoDB")
	return nil
}

package pending_execution

import (
	"context"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DdbTracker struct {
	client *dynamodb.Client
	table  string
}

func NewDdbTracker(table string, client *dynamodb.Client) Tracker {
	return &DdbTracker{
		table:  table,
		client: client,
	}
}

func (t *DdbTracker) AddExecution(ctx context.Context, id string, user tempest.Snowflake) error {
	log := logger.FromContext(ctx)
	log.Debug().Str("id", id).Str("user", user.String()).Msg("Adding execution")

	_, err := t.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &t.table,
		Item: map[string]types.AttributeValue{
			"id":   &types.AttributeValueMemberS{Value: id},
			"user": &types.AttributeValueMemberN{Value: user.String()},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to add execution")
		return err
	}

	return nil
}

func (t *DdbTracker) RemoveExecution(ctx context.Context, id string) error {
	log := logger.FromContext(ctx)
	log.Debug().Str("id", id).Msg("Removing execution")

	_, err := t.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &t.table,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to remove execution")
		return err
	}

	return nil
}

func (t *DdbTracker) GetExecution(ctx context.Context, id string) (tempest.Snowflake, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting execution")

	res, err := t.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &t.table,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to get execution")
		return 0, err
	}

	if res.Item == nil {
		log.Warn().Str("id", id).Msg("Execution not found")
		return 0, ErrExecutionNotFound
	}

	user := res.Item["user"].(*types.AttributeValueMemberN).Value
	snowlflake, err := tempest.StringToSnowflake(user)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert user ID")
		return 0, err
	}
	return snowlflake, nil
}

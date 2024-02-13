package dynamo_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (d *DynamoNodeRepository) Delete(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting node from DynamoDB")

	log.Debug().Msg("Calling DynamoDB")
	_, err := d.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &d.table,
		Key: map[string]types.AttributeValue{
			AttributeIdentifier: &types.AttributeValueMemberS{Value: id.String()},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Error deleting node from DynamoDB")
		return err
	}

	log.Debug().Msg("Deleted node from DynamoDB")
	return nil
}

package dynamo_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (d *DynamoNodeRepository) Create(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node in DynamoDB")

	log.Debug().Msg("Marshalling node to DynamoDB")
	record := recordFromNode(node)
	av, err := attributevalue.MarshalMap(record)
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling node to DynamoDB")
		return err
	}

	log.Debug().Msg("Calling DynamoDB")
	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &d.table,
		Item:      av,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error creating node in DynamoDB")
		return err
	}

	log.Debug().Msg("Created node in DynamoDB")
	return nil
}

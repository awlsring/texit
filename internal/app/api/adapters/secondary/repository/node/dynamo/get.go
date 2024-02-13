package dynamo_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (d *DynamoNodeRepository) Get(ctx context.Context, id node.Identifier) (*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting node from DynamoDB")

	log.Debug().Msg("Getting node from DynamoDB")
	resp, err := d.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: &d.table,
		Key: map[string]types.AttributeValue{
			AttributeIdentifier: &types.AttributeValueMemberS{Value: id.String()},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Error getting node from DynamoDB")
		return nil, err
	}
	if len(resp.Item) == 0 {
		log.Warn().Msg("Node not found")
		return nil, repository.ErrNodeNotFound
	}

	log.Debug().Msg("Got node from DynamoDB")
	node, err := recordFromDdb(resp.Item)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling node from DynamoDB")
		return nil, err
	}

	log.Debug().Msg("Got node from DynamoDB")
	return node.ToNode(), nil
}

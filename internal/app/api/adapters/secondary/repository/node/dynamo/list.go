package dynamo_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (d *DynamoNodeRepository) List(ctx context.Context) ([]*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Listing nodes from DynamoDB")

	log.Debug().Msg("Scanning DynamoDB table")
	resp, err := d.client.Scan(ctx, &dynamodb.ScanInput{
		TableName: &d.table,
	})
	if err != nil {
		log.Error().Err(err).Msg("Error listing nodes from DynamoDB")
		return nil, err
	}

	log.Debug().Msg("Scanned DynamoDB table")
	nodes := make([]*node.Node, 0, len(resp.Items))
	for _, item := range resp.Items {
		node, err := recordFromDdb(item)
		if err != nil {
			log.Error().Err(err).Msg("Error unmarshalling node from DynamoDB")
			return nil, err
		}
		nodes = append(nodes, node.ToNode())
	}

	log.Debug().Msg("Listed nodes from DynamoDB")
	return nodes, nil
}

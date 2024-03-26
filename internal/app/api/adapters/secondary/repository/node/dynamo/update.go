package dynamo_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func keyFromNode(node *node.Node) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{
		AttributeIdentifier: &types.AttributeValueMemberS{Value: node.Identifier.String()},
	}
}

func (d *DynamoNodeRepository) Update(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Updating node in DynamoDB")

	log.Debug().Msg("Creating update expression")
	expres := expression.Set(expression.Name("UpdatedAt"), expression.Value(node.UpdatedAt))
	expres = expres.Set(expression.Name("ProvisionStatus"), expression.Value(node.ProvisionStatus))
	if node.PlatformIdentifier.String() != "" {
		expres = expres.Set(expression.Name("PlatformIdentifier"), expression.Value(node.PlatformIdentifier))
	}
	if node.Provider.String() != "" {
		expres = expres.Set(expression.Name("Provider"), expression.Value(node.Provider))
	}
	if node.Location.String() != "" {
		expres = expres.Set(expression.Name("Location"), expression.Value(node.Location))
	}
	if node.PreauthKey.String() != "" {
		expres = expres.Set(expression.Name("PreauthKey"), expression.Value(node.PreauthKey))
	}
	if node.Tailnet.String() != "" {
		expres = expres.Set(expression.Name("Tailnet"), expression.Value(node.Tailnet))
	}
	if node.TailnetName.String() != "" {
		expres = expres.Set(expression.Name("TailnetName"), expression.Value(node.TailnetName))
	}
	if node.TailnetIdentifier.String() != "" {
		expres = expres.Set(expression.Name("TailnetIdentifier"), expression.Value(node.TailnetIdentifier))
	}
	if node.Size.String() != "" {
		expres = expres.Set(expression.Name("Size"), expression.Value(node.Size))
	}
	if node.Ephemeral {
		expres = expres.Set(expression.Name("Ephemeral"), expression.Value(node.Ephemeral))
	}

	log.Debug().Msg("Marshalling update expression")
	update, err := expression.NewBuilder().WithUpdate(expres).Build()
	if err != nil {
		log.Error().Err(err).Msg("Error marshalling update expression")
		return err
	}

	log.Debug().Msg("Calling DynamoDB")
	_, err = d.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName:                 &d.table,
		Key:                       keyFromNode(node),
		UpdateExpression:          update.Update(),
		ExpressionAttributeNames:  update.Names(),
		ExpressionAttributeValues: update.Values(),
	})

	if err != nil {
		log.Error().Err(err).Msg("Error updating node in DynamoDB")
		return err
	}

	log.Debug().Msg("Updated node in DynamoDB")
	return nil
}

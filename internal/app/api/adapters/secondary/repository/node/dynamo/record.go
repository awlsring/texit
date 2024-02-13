package dynamo_node_repository

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/rs/zerolog/log"
)

const (
	AttributeIdentifier         = "identifier"
	AttributePlatformIdentifier = "platform_identifier"
	AttributeProvider           = "provider_identifier"
	AttributeTailnet            = "tailnet"
	AttributeTailnetName        = "tailnet_device_name"
	AttributeTailnetIdentifier  = "tailnet_identifier"
	AttributeLocation           = "location"
	AttributePreauthKey         = "preauth_key"
	AttributeEphemeral          = "ephemeral"
	AttributeCreatedAt          = "created_at"
	AttributeUpdatedAt          = "updated_at"
)

type NodeDdbRecord struct {
	Identifier         string    `dynamodbav:"identifier"`
	PlatformIdentifier string    `dynamodbav:"platform_identifier"`
	Provider           string    `dynamodbav:"provider_identifier"`
	Tailnet            string    `dynamodbav:"tailnet"`
	TailnetName        string    `dynamodbav:"tailnet_device_name"`
	TailnetIdentifier  string    `dynamodbav:"tailnet_identifier"`
	Location           string    `dynamodbav:"location"`
	PreauthKey         string    `dynamodbav:"preauth_key"`
	Ephemeral          bool      `dynamodbav:"ephemeral"`
	CreatedAt          time.Time `dynamodbav:"created_at"`
	UpdatedAt          time.Time `dynamodbav:"updated_at"`
}

func recordFromNode(e *node.Node) *NodeDdbRecord {
	return &NodeDdbRecord{
		Identifier:         e.Identifier.String(),
		PlatformIdentifier: e.PlatformIdentifier.String(),
		Provider:           e.Provider.String(),
		Tailnet:            e.Tailnet.String(),
		TailnetName:        e.TailnetName.String(),
		TailnetIdentifier:  e.TailnetIdentifier.String(),
		Location:           e.Location.String(),
		PreauthKey:         e.PreauthKey.String(),
		Ephemeral:          e.Ephemeral,
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
	}
}

func recordFromDdb(item map[string]types.AttributeValue) (*NodeDdbRecord, error) {
	var entry NodeDdbRecord
	log.Debug().Msg("Unmarshalling execution from DynamoDB")
	err := attributevalue.UnmarshalMap(item, &entry)
	if err != nil {
		log.Error().Err(err).Msg("Error unmarshalling execution from DynamoDB")
		return nil, err
	}
	log.Debug().Msg("Unmarshalled execution from DynamoDB")
	return &entry, nil
}

func (n *NodeDdbRecord) ToNode() *node.Node {
	return &node.Node{
		Identifier:         node.Identifier(n.Identifier),
		PlatformIdentifier: node.PlatformIdentifier(n.PlatformIdentifier),
		Provider:           provider.Identifier(n.Provider),
		Tailnet:            tailnet.Identifier(n.Tailnet),
		TailnetIdentifier:  tailnet.DeviceIdentifier(n.TailnetIdentifier),
		TailnetName:        tailnet.DeviceName(n.TailnetName),
		Location:           provider.Location(n.Location),
		PreauthKey:         tailnet.PreauthKey(n.PreauthKey),
		Ephemeral:          n.Ephemeral,
		CreatedAt:          n.CreatedAt,
		UpdatedAt:          n.UpdatedAt,
	}
}

package config

import (
	"errors"

	"github.com/awlsring/texit/internal/pkg/config"
)

type TrackerType string

const (
	TrackerTypeInMemory TrackerType = "memory"
	TrackerTypeDynamoDb TrackerType = "ddb"
)

const (
	TrackerAwsAccessKey = "TRACKER_AWS_ACCESS_KEY_ID"
	TrackerAwsSecretKey = "TRACKER_AWS_SECRET_ACCESS_KEY"
	TrackerAwsRegion    = "TRACKER_AWS_REGION"
)

type TrackerConfig struct {
	Type      TrackerType `yaml:"type"`
	AccessKey string      `yaml:"accessKey"`
	SecretKey string      `yaml:"secretKey"`
	Region    string      `yaml:"region"`
}

func NewDefaultTrackerConfig() *TrackerConfig {
	return &TrackerConfig{
		Type: TrackerTypeInMemory,
	}
}

func (c *TrackerConfig) Validate() error {
	switch c.Type {
	case TrackerTypeInMemory:
		return nil
	case TrackerTypeDynamoDb:
		return c.ValidateDynamoDb()
	default:
		return errors.New("unknown notifier type")
	}
}

func (c *TrackerConfig) ValidateDynamoDb() error {
	if c.AccessKey == "" {
		val, err := config.AwsAccessKeyFromEnv(TrackerAwsAccessKey)
		if err == nil {
			c.AccessKey = val
		}
	}

	if c.SecretKey == "" {
		val, err := config.SecretKeyFromEnv(TrackerAwsSecretKey)
		if err == nil {
			c.SecretKey = val
		}
	}

	if c.Region == "" {
		val, err := config.RegionFromEnv(TrackerAwsRegion)
		if err != nil {
			return err
		}
		c.Region = val
	}

	return nil
}

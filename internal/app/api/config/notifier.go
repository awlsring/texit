package config

import (
	"errors"
	"os"

	"github.com/awlsring/texit/internal/pkg/config"
)

type NotifierType string

const (
	NotifierTypeMqtt NotifierType = "mqtt"
	NotifierTypeSns  NotifierType = "sns"
)

var (
	ErrMissingNotifierBroker   = errors.New("missing notifier broker")
	ErrMissingNotifierTopic    = errors.New("missing notifier topic")
	ErrMissiningNotifierRegion = errors.New("missing notifier region")
)

const (
	SnsAccessKeyEnvVar = "SNS_AWS_ACCESS_KEY_ID"
	SnsSecretKeyEnvVar = "SNS_AWS_SECRET_ACCESS_KEY"
	SnsRegionEnvVar    = "SNS_AWS_REGION"
)

// Configuration for the notifier
type NotifierConfig struct {
	// the notifier type
	Type NotifierType `yaml:"type"`
	// the topic to publish to
	Topic string `yaml:"topic"`
	// username for the notifier
	Username string `yaml:"username"`
	// password for the notifier
	Password string `yaml:"password"`
	// Access key for the notifier
	AccessKey string `yaml:"accessKey"`
	// Secret key for the notifier
	SecretKey string `yaml:"secretKey"`
	// the aws region for the notifier
	Region string `yaml:"region"`
	// The broker for the notifier
	Broker string `yaml:"broker"`
}

func (c *NotifierConfig) Validate() error {
	switch c.Type {
	case NotifierTypeMqtt:
		return c.validateMqtt()
	case NotifierTypeSns:
		return c.validateSns()
	default:
		return errors.New("unknown notifier type")
	}
}

func (c *NotifierConfig) validateSns() error {
	if c.Topic == "" {
		topic := os.Getenv("SNS_NOTIFIER_ARN")
		if topic == "" {
			return ErrMissingNotifierTopic
		}
		c.Topic = topic
	}
	if c.Region == "" {
		val, err := config.RegionFromEnv(SnsRegionEnvVar)
		if err != nil {
			return err
		}
		c.Region = val
	}
	if c.AccessKey == "" {
		val, err := config.AwsAccessKeyFromEnv(SnsAccessKeyEnvVar)
		if err == nil {
			c.AccessKey = val
		}
	}

	if c.SecretKey == "" {
		val, err := config.AwsAccessKeyFromEnv(SnsSecretKeyEnvVar)
		if err == nil {
			c.AccessKey = val
		}
	}
	return nil
}

func (c *NotifierConfig) validateMqtt() error {
	if c.Broker == "" {
		return ErrMissingNotifierBroker
	}
	if c.Topic == "" {
		return ErrMissingNotifierTopic
	}
	return nil
}

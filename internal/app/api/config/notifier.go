package config

import (
	"errors"
	"os"
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
		return ErrMissiningNotifierRegion
	}
	if c.AccessKey == "" {
		key := os.Getenv("SNS_AWS_ACCESS_KEY_ID")
		if key == "" {
			c.AccessKey = key
		}
	}

	if c.SecretKey == "" {
		key := os.Getenv("SNS_AWS_SECRET_ACCESS_KEY")
		if key != "" {
			c.SecretKey = key
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

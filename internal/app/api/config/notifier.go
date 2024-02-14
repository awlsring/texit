package config

import "errors"

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
	// the endpoint to publish to
	Endpoint string `yaml:"endpoint"`
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
		return ErrMissingNotifierTopic
	}
	if c.Region == "" {
		return ErrMissiningNotifierRegion
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

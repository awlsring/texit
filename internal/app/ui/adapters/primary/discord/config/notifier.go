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

type NotifierConfig struct {
	Broker   string       `yaml:"broker"`
	Topic    string       `yaml:"topic"`
	Type     NotifierType `yaml:"type"`
	User     string       `yaml:"user"`
	Password string       `yaml:"password"`
}

func (c *NotifierConfig) Validate() error {
	switch c.Type {
	case NotifierTypeMqtt:
		return c.validateMqtt()
	case NotifierTypeSns:
		return nil
	default:
		return errors.New("unknown notifier type")
	}
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

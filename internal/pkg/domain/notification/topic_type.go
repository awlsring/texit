package notification

import (
	"errors"
	"strings"
)

var (
	ErrUnknownTopicType = errors.New("unknown topic type")
)

type TopicType int

const (
	TopicTypeUnknown TopicType = iota
	TopicTypeSns
	TopicTypeMqtt
)

func (x TopicType) String() string {
	switch x {
	case TopicTypeSns:
		return "sns"
	case TopicTypeMqtt:
		return "mqtt"
	default:
		return "unknown"
	}
}

func TopicTypeFromString(s string) (TopicType, error) {
	switch strings.ToLower(s) {
	case "sns":
		return TopicTypeSns, nil
	case "mqtt":
		return TopicTypeMqtt, nil
	default:
		return TopicTypeUnknown, ErrUnknownTopicType
	}
}

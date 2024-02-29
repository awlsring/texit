package conversion

import (
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func TranslateNotifierrType(t notification.TopicType) texit.NotifierType {
	switch t {
	case notification.TopicTypeMqtt:
		return texit.NotifierTypeMqtt
	case notification.TopicTypeSns:
		return texit.NotifierTypeAWSSns
	default:
		return texit.NotifierTypeUnknown
	}
}

func NotifierToSummary(n *notification.Notifier) texit.NotifierSummary {
	return texit.NotifierSummary{
		Name:     n.Name.String(),
		Endpoint: n.Endpoint.String(),
		Type:     TranslateNotifierrType(n.Type),
	}
}

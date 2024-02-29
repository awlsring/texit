package notification

type Notifier struct {
	Name     Identifier
	Type     TopicType
	Endpoint Endpoint
}

func NewNotifier(name Identifier, topicType TopicType, endpoint Endpoint) *Notifier {
	return &Notifier{
		Name:     name,
		Type:     topicType,
		Endpoint: endpoint,
	}
}

func NewNotifierFromPrimitive(name string, topicType string, endpoint string) (*Notifier, error) {
	n, err := IdentifierFromString(name)
	if err != nil {
		return nil, err
	}
	t, err := TopicTypeFromString(topicType)
	if err != nil {
		return nil, err
	}
	e, err := EndpointFromString(endpoint)
	if err != nil {
		return nil, err
	}
	return NewNotifier(n, t, e), nil
}

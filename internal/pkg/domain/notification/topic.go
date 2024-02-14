package notification

type Topic struct {
	TopicType     TopicType
	TopicEndpoint TopicEndpoint
}

func NewTopic(topicType TopicType, topicEndpoint TopicEndpoint) Topic {
	return Topic{
		TopicType:     topicType,
		TopicEndpoint: topicEndpoint,
	}
}

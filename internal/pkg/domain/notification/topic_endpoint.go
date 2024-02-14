package notification

import "github.com/awlsring/texit/internal/pkg/values"

type TopicEndpoint string

func (x TopicEndpoint) String() string {
	return string(x)
}

func TopicEndpointFromString(s string) (TopicEndpoint, error) {
	end, err := values.NonNullString[TopicEndpoint](s)
	if err != nil {
		return "", err
	}
	return TopicEndpoint(end), nil
}

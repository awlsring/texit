package notification

import "github.com/awlsring/texit/internal/pkg/values"

type Endpoint string

func (i Endpoint) String() string {
	return string(i)
}

func EndpointFromString(e string) (Endpoint, error) {
	end, err := values.NonNullString[Endpoint](e)
	if err != nil {
		return "", err
	}
	return Endpoint(end), nil
}

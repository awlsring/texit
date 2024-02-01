package tailnet

import "github.com/awlsring/texit/internal/pkg/values"

type ControlServer string

func (i ControlServer) String() string {
	return string(i)
}

func ControlServerFromString(u string) (ControlServer, error) {
	c, err := values.NonNullString[ControlServer](u)
	if err != nil {
		return "", err
	}
	return ControlServer(c), nil
}

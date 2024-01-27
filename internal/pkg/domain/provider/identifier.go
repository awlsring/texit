package provider

import "github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/values"

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func IdentifierFromString(id string) (Identifier, error) {
	identifier, err := values.NonNullString[Identifier](id)
	if err != nil {
		return "", err
	}
	return Identifier(identifier), nil
}

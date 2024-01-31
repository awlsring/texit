package tailnet

import (
	"github.com/awlsring/texit/internal/pkg/values"
	"github.com/pkg/errors"
)

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func IdentifierFromString(s string) (Identifier, error) {
	id, err := values.NonNullString[Identifier](s)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse tailnet identifier")
	}
	return Identifier(id), nil
}

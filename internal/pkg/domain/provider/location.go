package provider

import (
	"github.com/awlsring/texit/internal/pkg/values"
	"github.com/pkg/errors"
)

type Location string

func (l Location) String() string {
	return string(l)
}

var (
	ErrUnknownLocation = errors.New("unknown location")
)

func LocationFromString(location string) (Location, error) {
	l, err := values.NonNullString[Identifier](location)
	if err != nil {
		return "", err
	}
	return Location(l), nil
}

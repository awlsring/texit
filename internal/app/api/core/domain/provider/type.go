package provider

import (
	"strings"

	"github.com/pkg/errors"
)

type Type int

const (
	TypeAwsEcs Type = iota
	TypeUnknown
)

var (
	ErrUnknownType = errors.New("unknown type")
)

func (t Type) String() string {
	switch t {
	case TypeAwsEcs:
		return "aws-ecs"
	default:
		return "unknown"
	}
}

func TypeFromString(s string) (Type, error) {
	switch strings.ToLower(s) {
	case "aws-ecs":
		return TypeAwsEcs, nil
	default:
		return TypeUnknown, errors.Wrap(ErrUnknownType, s)
	}
}

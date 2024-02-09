package provider

import (
	"strings"

	"github.com/pkg/errors"
)

type Type int

const (
	TypeUnknown Type = iota
	TypeAwsEcs
	TypeAwsEc2
	TypeLinode
)

var (
	ErrUnknownType = errors.New("unknown type")
)

func (t Type) String() string {
	switch t {
	case TypeAwsEcs:
		return "aws-ecs"
	case TypeAwsEc2:
		return "aws-ec2"
	case TypeLinode:
		return "linode"
	default:
		return "unknown"
	}
}

func TypeFromString(s string) (Type, error) {
	switch strings.ToLower(s) {
	case "aws-ecs":
		return TypeAwsEcs, nil
	case "aws-ec2":
		return TypeAwsEc2, nil
	case "linode":
		return TypeLinode, nil
	default:
		return TypeUnknown, errors.Wrap(ErrUnknownType, s)
	}
}

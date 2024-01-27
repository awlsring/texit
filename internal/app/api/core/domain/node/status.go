package node

import "github.com/pkg/errors"

type Status int

const (
	StatusUnknown Status = iota
	StatusActive
	StatusInactive
)

var (
	ErrUnknownStatus = errors.New("unknown status")
)

func (s Status) String() string {
	switch s {
	case StatusActive:
		return "active"
	case StatusInactive:
		return "inactive"
	default:
		return "unknown"
	}
}

func StatusFromString(s string) (Status, error) {
	switch s {
	case "active":
		return StatusActive, nil
	case "inactive":
		return StatusInactive, nil
	default:
		return StatusUnknown, errors.Wrap(ErrUnknownStatus, s)
	}
}

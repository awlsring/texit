package workflow

import (
	"strings"

	"github.com/pkg/errors"
)

type Status int

const (
	StatusUnknown Status = iota
	StatusPending
	StatusRunning
	StatusComplete
	StatusFailed
)

func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "unknown"
	case StatusPending:
		return "pending"
	case StatusRunning:
		return "running"
	case StatusComplete:
		return "complete"
	case StatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

func StatusFromString(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "pending":
		return StatusPending, nil
	case "running":
		return StatusRunning, nil
	case "complete":
		return StatusComplete, nil
	case "failed":
		return StatusFailed, nil
	default:
		return StatusUnknown, errors.Wrap(ErrUnknownStatus, s)
	}
}

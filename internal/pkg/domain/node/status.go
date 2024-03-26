package node

import (
	"strings"

	"github.com/pkg/errors"
)

type Status int

const (
	StatusUnknown Status = iota
	StatusRunning
	StatusStarting
	StatusStopping
	StatusStopped
	StatusPending
)

var (
	ErrUnknownStatus = errors.New("unknown status")
)

func (s Status) String() string {
	switch s {
	case StatusRunning:
		return "running"
	case StatusStarting:
		return "starting"
	case StatusStopping:
		return "stopping"
	case StatusStopped:
		return "stopped"
	case StatusPending:
		return "pending"
	default:
		return "unknown"
	}
}

func StatusFromString(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "running":
		return StatusRunning, nil
	case "starting":
		return StatusStarting, nil
	case "stopping":
		return StatusStopping, nil
	case "stopped":
		return StatusStopped, nil
	case "pending":
		return StatusPending, nil
	default:
		return StatusUnknown, ErrUnknownStatus
	}
}

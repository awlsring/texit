package node

import (
	"strings"

	"github.com/pkg/errors"
)

type ProvisionStatus int

const (
	ProvisionStatusUnknown ProvisionStatus = iota
	ProvisionStatusCreating
	ProvisionStatusCreated
	ProvisionStatusFailed
)

var (
	ErrUnknownProvisionStatus = errors.New("unknown provision status")
)

func (s ProvisionStatus) String() string {
	switch s {
	case ProvisionStatusCreating:
		return "creating"
	case ProvisionStatusCreated:
		return "created"
	case ProvisionStatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

func ProvisionStatusFromString(s string) (ProvisionStatus, error) {
	switch strings.ToLower(s) {
	case "creating":
		return ProvisionStatusCreating, nil
	case "created":
		return ProvisionStatusCreated, nil
	case "failed":
		return ProvisionStatusFailed, nil
	default:
		return ProvisionStatusUnknown, ErrUnknownProvisionStatus
	}
}

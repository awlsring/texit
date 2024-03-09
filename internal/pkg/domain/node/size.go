package node

import (
	"strings"

	"github.com/pkg/errors"
)

type Size int

const (
	SizeUnknown Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
)

var (
	ErrUnknownSize = errors.New("unknown size")
)

func (s Size) String() string {
	switch s {
	case SizeSmall:
		return "small"
	case SizeMedium:
		return "medium"
	case SizeLarge:
		return "large"
	default:
		return "unknown"
	}
}

func SizeFromString(s string) (Size, error) {
	switch strings.ToLower(s) {
	case "small":
		return SizeSmall, nil
	case "medium":
		return SizeMedium, nil
	case "large":
		return SizeLarge, nil
	default:
		return SizeUnknown, ErrUnknownSize
	}
}

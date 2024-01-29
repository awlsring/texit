package tailnet

import "fmt"

type Type int

const (
	TypeUnknown Type = iota
	TypeTailscale
	TypeHeadscale
)

func (t Type) String() string {
	switch t {
	case TypeTailscale:
		return "tailscale"
	case TypeHeadscale:
		return "headscale"
	default:
		return "unknown"
	}
}

func TypeFromString(s string) (Type, error) {
	switch s {
	case "tailscale":
		return TypeTailscale, nil
	case "headscale":
		return TypeHeadscale, nil
	default:
		return TypeUnknown, fmt.Errorf("invalid tailnet type: %s", s)
	}
}

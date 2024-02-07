package tsn

import "github.com/pkg/errors"

var (
	ErrMissingTailnetAuthKey       = errors.New("missing tailnet preauth key")
	ErrMissingHostname             = errors.New("missing hostname")
	ErrInvalidServerConfigurations = errors.New("invalid server configurations")
	ErrUnknownServerMode           = errors.New("unknown server mode")
)

type ServerMode int8

const (
	ServerModeFunnel ServerMode = iota
	ServerModeTls
	SeverModeStandard
	ServerModeUnknown
)

func (s ServerMode) String() string {
	switch s {
	case ServerModeFunnel:
		return "funnel"
	case ServerModeTls:
		return "tls"
	default:
		return "standard"
	}
}

func ServerModeFromString(s string) (ServerMode, error) {
	switch s {
	case "funnel":
		return ServerModeFunnel, nil
	case "tls":
		return ServerModeTls, nil
	case "standard":
		return SeverModeStandard, nil
	default:
		return ServerModeUnknown, errors.Wrap(ErrUnknownServerMode, s)
	}
}

type Config struct {
	// Required, the authkey to join the tailnet
	AuthKey string `yaml:"authkey"`
	// The hostname to use on the tailnet
	Hostname string `yaml:"hostname"`
	// The directory to store the state of the tailnet. If not specified, the default will be used.
	StateDir string `yaml:"state"`
	// The mode to start the server, can be standard, tls or funnel. Defaults to standard.
	ModeStr string `yaml:"mode"`
	// internal translation of string mode
	mode ServerMode
	// ControlUrl is the URL of the control server to use. Specify this if you are using Headscale. If not specified, the default tailscale address will be used.
	ControlUrl string `yaml:"controlUrl"`
}

func (c *Config) Mode() ServerMode {
	return c.mode
}

func (c *Config) Validate() error {
	if c.AuthKey == "" {
		return ErrMissingTailnetAuthKey
	}

	if c.Hostname == "" {
		return ErrMissingHostname
	}

	if c.ModeStr == "" {
		c.mode = SeverModeStandard
	} else {
		m, err := ServerModeFromString(c.ModeStr)
		if err != nil {
			return err
		}
		c.mode = m
	}

	return nil
}

package config

import "errors"

type TailnetType string // tailscale or headscale
const (
	TailnetTypeTailscale TailnetType = "tailscale"
	TailnetTypeHeadscale TailnetType = "headscale"
)

func (t TailnetType) String() string {
	return string(t)
}

var (
	ErrMissingTailnetType            = errors.New("missing tailnet type")
	ErrMissingTailnet                = errors.New("missing tailnet")
	ErrMissingTailnetApiKey          = errors.New("missing tailnet api key")
	ErrMissingUser                   = errors.New("missing headscale user")
	ErrMissingHeadscaleControlServer = errors.New("missing headscale control server")
)

// Configuration for the tailnet exit nodes will join
type TailnetConfig struct {
	// The type of tailnet, tailscale or headscale
	Type TailnetType `yaml:"type"`
	// The network of the tailnet. On tailscale this is the network id
	Tailnet string `yaml:"tailnet"`
	// The api token to communicate with the tailnet
	ApiKey string `yaml:"apiKey"`
	// The user to register exist nodes for.
	User string `yaml:"user"`
	// the control server to use. This is require for headscale
	ControlServer string `yaml:"controlServer"`
}

func (c *TailnetConfig) Validate() error {
	if c.Type == "" {
		return ErrMissingTailnetType
	}
	if c.Tailnet == "" {
		return ErrMissingTailnet
	}
	if c.ApiKey == "" {
		return ErrMissingTailnetApiKey
	}

	if c.User == "" {
		return ErrMissingUser
	}

	if c.Type == TailnetTypeHeadscale && c.ControlServer == "" {
		return errors.New("missing control server")
	}

	if c.ControlServer == "" {
		c.ControlServer = "https://controlplane.tailscale.com"
	}

	return nil
}

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
	ErrMissingTailnetType   = errors.New("missing tailnet type")
	ErrMissingTailnet       = errors.New("missing tailnet")
	ErrMissingTailnetApiKey = errors.New("missing tailnet api key")
	ErrMissingHeadscaleUser = errors.New("missing headscale user")
)

// Configuration for the tailnet exit nodes will join
type TailnetConfig struct {
	// The type of tailnet, tailscale or headscale
	Type TailnetType `yaml:"type"`
	// The network of the tailnet. On tailscale, this is your tailnet name. On headscale, this is the server address.
	Tailnet string `yaml:"tailnet"`
	// The api token to communicate with the tailnet
	ApiKey string `yaml:"apiKey"`
	// The user to register exist nodes for.
	User string
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
		return ErrMissingHeadscaleUser
	}

	return nil
}

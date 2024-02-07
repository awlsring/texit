package tsn

import (
	"net"

	"github.com/pkg/errors"
)

const (
	network = "tcp"
)

var (
	ErrInvalidAddress = errors.New("invalid address")
	ValidFunnelAddr   = []string{":443", ":8443", ":1000"}
)

func validFunnelAddr(addr string) bool {
	for _, a := range ValidFunnelAddr {
		if a == addr {
			return true
		}
	}
	return false
}

func ListenerFromConfig(cfg Config, addr string, opts ...ServerOption) (net.Listener, error) {
	bopts := []ServerOption{
		WithAuthKey(cfg.AuthKey),
		WithHostname(cfg.Hostname),
	}

	if cfg.StateDir != "" {
		opts = append(opts, WithStateDir(cfg.StateDir))
	}

	if cfg.ControlUrl != "" {
		opts = append(opts, WithControlURL(cfg.ControlUrl))
	}

	srv := NewServer(bopts...)

	for _, opt := range opts {
		opt(srv)
	}

	switch cfg.Mode() {
	case ServerModeFunnel:
		if !validFunnelAddr(addr) {
			return nil, errors.Wrap(ErrInvalidAddress, addr)
		}
		return srv.ListenFunnel(network, addr)
	case ServerModeTls:
		return srv.ListenTLS(network, addr)
	default:
		return srv.Listen(network, addr)
	}
}

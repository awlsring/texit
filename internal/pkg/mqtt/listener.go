package mqtt

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ListenHandler interface {
	Handle(context.Context, mqtt.Message)
}

type Listener interface {
	Subscribe(context.Context, string) error
}

type listener struct {
	lvl zerolog.Level
	log zerolog.Logger
	hdl ListenHandler
	m   mqtt.Client
}

func (l *listener) publishHandler(client mqtt.Client, msg mqtt.Message) {
	ctx := logger.InitContextLogger(context.Background(), l.lvl)
	l.hdl.Handle(ctx, msg)
}

func (l *listener) connectHandler(mqtt.Client) {
	log.Info().Msg("Connected to broker")
}

func (l *listener) connectDropHandler(mqtt.Client, error) {
	log.Error().Msg("Connection dropped")
}

func (l *listener) Subscribe(ctx context.Context, t string) error {
	if token := l.m.Subscribe(t, 0, nil); token.Wait() && token.Error() != nil {
		log.Error().Err(token.Error()).Msg("Failed to subscribe")
		return token.Error()
	}
	return nil
}

func NewListener(addr string, hdl ListenHandler, opts ...ListenerOption) (Listener, error) {
	l := &listener{
		lvl: zerolog.InfoLevel,
		hdl: hdl,
	}

	for _, o := range opts {
		o(l)
	}

	o := mqtt.NewClientOptions()
	o.AddBroker(addr)
	o.SetDefaultPublishHandler(l.publishHandler)
	o.SetOnConnectHandler(l.connectHandler)
	o.SetConnectionLostHandler(l.connectDropHandler)
	o.SetAutoReconnect(true)
	c := mqtt.NewClient(o)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	l.m = c

	return l, nil
}

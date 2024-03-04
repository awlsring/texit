package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type Class int

const (
	Counter Class = iota
	UpDownCounter
	Histogram
)

type Metrics interface {
	Counter(ctx context.Context, met string, value float64, opt ...metric.AddOption) error
	UpDownCounter(ctx context.Context, met string, value float64, opt ...metric.AddOption) error
	Histogram(ctx context.Context, met string, value float64, opt ...metric.RecordOption) error
}

type MetricsOpts func(*metrics)

func WithNamespace(namespace string) MetricsOpts {
	return func(m *metrics) {
		m.namespace = namespace
	}
}

type metrics struct {
	namespace string
	tracer    trace.Tracer
	meter     metric.Meter
	counters  map[string]metric.Float64Counter
	upDown    map[string]metric.Float64UpDownCounter
	histogram map[string]metric.Float64Histogram
}

func NewMetrics(opts ...MetricsOpts) Metrics {
	m := &metrics{
		namespace: "",
		counters:  make(map[string]metric.Float64Counter),
		upDown:    make(map[string]metric.Float64UpDownCounter),
		histogram: make(map[string]metric.Float64Histogram),
	}
	for _, opt := range opts {
		opt(m)
	}

	m.tracer = otel.Tracer(m.namespace)
	m.meter = otel.Meter(m.namespace)

	return m
}

func (m *metrics) Counter(ctx context.Context, met string, value float64, opt ...metric.AddOption) error {
	if _, ok := m.counters[met]; !ok {
		c, err := m.meter.Float64Counter(met)
		if err != nil {
			return err
		}
		m.counters[met] = c
	}
	m.counters[met].Add(ctx, value, opt...)

	return nil
}

func (m *metrics) UpDownCounter(ctx context.Context, met string, value float64, opt ...metric.AddOption) error {
	if _, ok := m.upDown[met]; !ok {
		c, err := m.meter.Float64UpDownCounter(met)
		if err != nil {
			return err
		}
		m.upDown[met] = c
	}
	m.upDown[met].Add(ctx, value, opt...)

	return nil
}

func (m *metrics) Histogram(ctx context.Context, met string, value float64, opt ...metric.RecordOption) error {
	if _, ok := m.histogram[met]; !ok {
		c, err := m.meter.Float64Histogram(met)
		if err != nil {
			return err
		}
		m.histogram[met] = c
	}
	m.histogram[met].Record(ctx, value, opt...)

	return nil
}

package observability

import (
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func NewPrometheusMetricsProvider(ex *prometheus.Exporter, opts ...metric.Option) *metric.MeterProvider {
	opts = append(opts, metric.WithReader(ex))
	meterProvider := metric.NewMeterProvider(opts...)
	return meterProvider
}

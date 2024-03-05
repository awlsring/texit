package observability

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewPrometheusMetricsProvider(ex *prometheus.Exporter, opts ...metric.Option) *metric.MeterProvider {
	re := resource.NewWithAttributes("", attribute.KeyValue{Key: "service.name", Value: attribute.StringValue("texit")})
	opts = append(opts, metric.WithReader(ex), metric.WithResource(re))
	meterProvider := metric.NewMeterProvider(opts...)
	return meterProvider
}

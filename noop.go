package gotel

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	otelmetric "go.opentelemetry.io/otel/metric"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// NoopTelemetry is a no-op implementation of the TelemetryProvider interface.
type NoopTelemetry struct {
	serviceName string
}

// NewNoopTelemetry creates a new NoopTelemetry instance.
func NewNoopTelemetry(cfg Config) (*NoopTelemetry, error) {
	return &NoopTelemetry{serviceName: cfg.ServiceName}, nil
}

// GetServiceName returns the service name.
func (t *NoopTelemetry) GetServiceName() string { return t.serviceName }

// LogInfo logs nothing.
func (t *NoopTelemetry) LogInfo(args ...interface{}) {}

// LogErrorln logs nothing.
func (t *NoopTelemetry) LogErrorln(args ...interface{}) {}

// LogFatalln logs nothing, then exits.
func (t *NoopTelemetry) LogFatalln(args ...interface{}) {
	os.Exit(1)
}

// LogRequest is a no-op middleware.
func (t *NoopTelemetry) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) { c.Next() }
}

// MeterRequestDuration is a no-op middleware.
func (t *NoopTelemetry) MeterRequestDuration() gin.HandlerFunc {
	return func(c *gin.Context) { c.Next() }
}

// MeterRequestsInFlight is a no-op middleware.
func (t *NoopTelemetry) MeterRequestsInFlight() gin.HandlerFunc {
	return func(c *gin.Context) { c.Next() }
}

// TraceStart returns the context and span unchanged.
func (t *NoopTelemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span) {
	return ctx, oteltrace.SpanFromContext(ctx)
}

// MeterInt64Histogram returns nil.
func (t *NoopTelemetry) MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error) {
	return nil, nil
}

// MeterInt64UpDownCounter returns nil.
func (t *NoopTelemetry) MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error) {
	return nil, nil
}

// Shutdown does nothing.
func (t *NoopTelemetry) Shutdown(ctx context.Context) {}

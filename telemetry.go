package gotel

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltrace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// TelemetryProvider is an interface for the telemetry provider.
type TelemetryProvider interface {
	GetServiceName() string
	LogInfo(args ...interface{})
	LogErrorln(args ...interface{})
	LogFatalln(args ...interface{})
	MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error)
	MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error)
	TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span)
	LogRequest() gin.HandlerFunc
	MeterRequestDuration() gin.HandlerFunc
	MeterRequestsInFlight() gin.HandlerFunc
	Shutdown(ctx context.Context)
}

// Telemetry is a wrapper around the OpenTelemetry logger, meter, and tracer.
type Telemetry struct {
	lp     *log.LoggerProvider
	mp     *metric.MeterProvider
	tp     *trace.TracerProvider
	log    *zap.SugaredLogger
	meter  otelmetric.Meter
	tracer oteltrace.Tracer
	cfg    Config
}

// NewTelemetry creates a new telemetry instance.
func NewTelemetry(ctx context.Context, cfg Config) (*Telemetry, error) {
	rp := newResource(cfg.ServiceName, cfg.ServiceVersion)

	lp, err := newLoggerProvider(ctx, rp)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger: %w", err)
	}

	logger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(os.Stdout), zapcore.InfoLevel),
			otelzap.NewCore(cfg.ServiceName, otelzap.WithLoggerProvider(lp)),
		),
	)

	mp, err := newMeterProvider(ctx, rp)
	if err != nil {
		return nil, fmt.Errorf("failed to create meter: %w", err)
	}
	meter := mp.Meter(cfg.ServiceName)

	tp, err := newTracerProvider(ctx, rp)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer: %w", err)
	}
	tracer := tp.Tracer(cfg.ServiceName)

	return &Telemetry{
		lp:     lp,
		mp:     mp,
		tp:     tp,
		log:    logger.Sugar(),
		meter:  meter,
		tracer: tracer,
		cfg:    cfg,
	}, nil
}

// GetServiceName returns the name of the service.
func (t *Telemetry) GetServiceName() string {
	return t.cfg.ServiceName
}

// LogInfo logs a message at the info level.
func (t *Telemetry) LogInfo(args ...interface{}) {
	t.log.Info(args...)
}

// LogErrorln logs a message and then calls os.Exit(1).
func (t *Telemetry) LogErrorln(args ...interface{}) {
	t.log.Errorln(args...)
}

// LogFatalln logs a message and then calls os.Exit(1).
func (t *Telemetry) LogFatalln(args ...interface{}) {
	t.log.Fatalln(args...)
}

// MeterInt64Histogram creates a new int64 histogram metric.
func (t *Telemetry) MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error) { //nolint:ireturn
	histogram, err := t.meter.Int64Histogram(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create histogram: %w", err)
	}

	return histogram, nil
}

// MeterInt64UpDownCounter creates a new int64 up down counter metric.
func (t *Telemetry) MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error) { //nolint:ireturn
	counter, err := t.meter.Int64UpDownCounter(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create counter: %w", err)
	}

	return counter, nil
}

// TraceStart starts a new span with the given name. The span must be ended by calling End.
func (t *Telemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltrace.Span) { //nolint:ireturn
	//nolint: spancheck
	return t.tracer.Start(ctx, name)
}

// Shutdown shuts down the logger, meter, and tracer.
func (t *Telemetry) Shutdown(ctx context.Context) {
	t.lp.Shutdown(ctx)
	t.mp.Shutdown(ctx)
	t.tp.Shutdown(ctx)
}

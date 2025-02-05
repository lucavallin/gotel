package gotel

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/semconv/v1.20.0/httpconv"
)

// LogRequest is a gin middleware that logs the request path.
func (t *Telemetry) LogRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		t.LogInfo("request to ", c.Request.URL.Path)
		c.Next()
		t.LogInfo("end of request to ", c.Request.URL.Path)
	}
}

// MeterRequestDuration is a gin middleware that captures the duration of the request.
func (t *Telemetry) MeterRequestDuration() gin.HandlerFunc {
	// init metric, here we are using histogram for capturing request duration
	histogram, err := t.MeterInt64Histogram(MetricRequestDurationMillis)
	if err != nil {
		t.LogFatalln(fmt.Errorf("failed to create histogram: %w", err))
	}

	return func(c *gin.Context) {
		// capture the start time of the request
		startTime := time.Now()

		// execute next http handler
		c.Next()

		// record the request duration
		duration := time.Since(startTime)
		histogram.Record(
			c.Request.Context(),
			duration.Milliseconds(),
			otelmetric.WithAttributes(
				httpconv.ServerRequest(t.GetServiceName(), c.Request)...,
			),
		)
	}
}

// MeterRequestsInFlight is a gin middleware that captures the number of requests in flight.
func (t *Telemetry) MeterRequestsInFlight() gin.HandlerFunc {
	// init metric, here we are using counter for capturing request in flight
	counter, err := t.MeterInt64UpDownCounter(MetricRequestsInFlight)
	if err != nil {
		t.LogFatalln(fmt.Errorf("failed to create counter: %w", err))
	}

	return func(c *gin.Context) {
		// define metric attributes
		attrs := otelmetric.WithAttributes(httpconv.ServerRequest(t.GetServiceName(), c.Request)...)

		// increase the number of requests in flight
		counter.Add(c.Request.Context(), 1, attrs)

		// execute next http handler
		c.Next()

		// decrease the number of requests in flight
		counter.Add(c.Request.Context(), -1, attrs)
	}
}

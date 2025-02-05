package gotel_test

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucavallin/gotel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewNoopTelemetry(t *testing.T) {
	cfg := gotel.Config{
		ServiceName:    "test-service",
		ServiceVersion: "1.0.0",
		Enabled:        false,
	}

	noop, err := gotel.NewNoopTelemetry(cfg)
	require.NoError(t, err)
	assert.NotNil(t, noop)

	// Test service name
	assert.Equal(t, cfg.ServiceName, noop.GetServiceName())

	// Test logging (should not panic)
	assert.NotPanics(t, func() {
		noop.LogInfo("test")
		noop.LogErrorln("test")
	})

	// Test metrics
	histogram, err := noop.MeterInt64Histogram(gotel.Metric{})
	require.NoError(t, err)
	assert.Nil(t, histogram)

	counter, err := noop.MeterInt64UpDownCounter(gotel.Metric{})
	require.NoError(t, err)
	assert.Nil(t, counter)

	// Test tracing
	ctx := context.Background()
	newCtx, span := noop.TraceStart(ctx, "test")
	assert.Equal(t, ctx, newCtx)
	assert.NotNil(t, span)

	// Test middleware functions
	assert.NotNil(t, noop.LogRequest())
	assert.NotNil(t, noop.MeterRequestDuration())
	assert.NotNil(t, noop.MeterRequestsInFlight())

	// Test middleware execution
	c, _ := gin.CreateTestContext(nil)
	noop.LogRequest()(c)
	noop.MeterRequestDuration()(c)
	noop.MeterRequestsInFlight()(c)
}

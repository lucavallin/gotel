package gotel_test

import (
	"context"
	"testing"

	"github.com/lucavallin/gotel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewTelemetry(t *testing.T) {
	tests := []struct {
		name    string
		cfg     gotel.Config
		wantErr bool
	}{
		{
			name: "success with telemetry enabled",
			cfg: gotel.Config{
				ServiceName:    "test-service",
				ServiceVersion: "1.0.0",
				Enabled:        true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			telem, err := gotel.NewTelemetry(ctx, tt.cfg)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, telem)

				return
			}

			require.NoError(t, err)
			assert.NotNil(t, telem)

			// Basic functionality test
			assert.Equal(t, tt.cfg.ServiceName, telem.GetServiceName())

			// Test metrics creation
			histogram, err := telem.MeterInt64Histogram(gotel.MetricRequestDurationMillis)
			if tt.cfg.Enabled {
				require.NoError(t, err)
				assert.NotNil(t, histogram)
			} else {
				assert.Nil(t, histogram)
			}

			counter, err := telem.MeterInt64UpDownCounter(gotel.MetricRequestsInFlight)
			if tt.cfg.Enabled {
				require.NoError(t, err)
				assert.NotNil(t, counter)
			} else {
				assert.Nil(t, counter)
			}

			// Clean up
			telem.Shutdown(ctx)
		})
	}
}

func TestTelemetry_MeterInt64Histogram(t *testing.T) {
	ctx := context.Background()
	telem, err := gotel.NewTelemetry(ctx, gotel.Config{
		ServiceName:    "test-service",
		ServiceVersion: "1.0.0",
		Enabled:        true,
	})
	require.NoError(t, err)
	defer telem.Shutdown(ctx)

	tests := []struct {
		name    string
		metric  gotel.Metric
		wantErr bool
	}{
		{
			name: "success with valid metric",
			metric: gotel.Metric{
				Name:        "test_histogram",
				Unit:        "ms",
				Description: "Test histogram",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			histogram, err := telem.MeterInt64Histogram(tt.metric)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, histogram)

				return
			}

			require.NoError(t, err)
			assert.NotNil(t, histogram)
		})
	}
}

func TestTelemetry_MeterInt64UpDownCounter(t *testing.T) {
	ctx := context.Background()
	telem, err := gotel.NewTelemetry(ctx, gotel.Config{
		ServiceName:    "test-service",
		ServiceVersion: "1.0.0",
	})
	require.NoError(t, err)
	defer telem.Shutdown(ctx)

	tests := []struct {
		name    string
		metric  gotel.Metric
		wantErr bool
	}{
		{
			name: "success with valid metric",
			metric: gotel.Metric{
				Name:        "test_counter",
				Unit:        "1",
				Description: "Test counter",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter, err := telem.MeterInt64UpDownCounter(tt.metric)
			if tt.wantErr {
				require.Error(t, err)
				assert.Nil(t, counter)

				return
			}

			require.NoError(t, err)
			assert.NotNil(t, counter)
		})
	}
}

func TestTelemetry_TraceStart(t *testing.T) {
	ctx := context.Background()
	telem, err := gotel.NewTelemetry(ctx, gotel.Config{
		ServiceName:    "test-service",
		ServiceVersion: "1.0.0",
		Enabled:        true,
	})
	require.NoError(t, err)
	defer telem.Shutdown(ctx)

	newCtx, span := telem.TraceStart(ctx, "test-span")
	assert.NotNil(t, newCtx)
	assert.NotNil(t, span)
	span.End()
}

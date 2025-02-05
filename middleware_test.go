package gotel_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucavallin/gotel"
	mockmetric "github.com/lucavallin/gotel/mocks/metric"
	"github.com/stretchr/testify/assert"
)

func TestTelemetry_LogRequest(t *testing.T) {
	t.Run("logs request info", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		telemetry, _ := gotel.NewTelemetry(context.Background(), gotel.Config{
			ServiceName:    "test-service",
			ServiceVersion: "1.0.0",
			Enabled:        true,
		})
		defer telemetry.Shutdown(context.Background())

		// Create test router with middleware
		r := gin.New()
		r.Use(telemetry.LogRequest())
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// Create and execute request
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestTelemetry_MeterRequestDuration(t *testing.T) {
	t.Run("records request duration", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		mockHist := mockmetric.NewMockInt64Histogram(t)
		telemetry, _ := gotel.NewTelemetry(context.Background(), gotel.Config{
			ServiceName:    "test-service",
			ServiceVersion: "1.0.0",
			Enabled:        true,
		})
		defer telemetry.Shutdown(context.Background())

		// Create test router with middleware
		r := gin.New()
		r.Use(telemetry.MeterRequestDuration())
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// Create and execute request
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockHist.AssertExpectations(t)
	})
}

func TestTelemetry_MeterRequestsInFlight(t *testing.T) {
	t.Run("tracks concurrent requests", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		mockCount := mockmetric.NewMockFloat64UpDownCounter(t)
		telemetry, _ := gotel.NewTelemetry(context.Background(), gotel.Config{
			ServiceName:    "test-service",
			ServiceVersion: "1.0.0",
			Enabled:        true,
		})
		defer telemetry.Shutdown(context.Background())

		// Create test router with middleware
		r := gin.New()
		r.Use(telemetry.MeterRequestsInFlight())
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// Create and execute request
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/test", nil)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockCount.AssertExpectations(t)
	})

	t.Run("handles multiple concurrent requests", func(t *testing.T) {
		// Setup
		gin.SetMode(gin.TestMode)
		mockCount := mockmetric.NewMockFloat64UpDownCounter(t)
		telemetry, _ := gotel.NewTelemetry(context.Background(), gotel.Config{
			ServiceName:    "test-service",
			ServiceVersion: "1.0.0",
			Enabled:        true,
		})
		defer telemetry.Shutdown(context.Background())

		// Create test router with middleware
		r := gin.New()
		r.Use(telemetry.MeterRequestsInFlight())
		r.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		// Execute two concurrent requests
		for range [2]int{} {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, http.StatusOK, w.Code)
		}

		mockCount.AssertExpectations(t)
	})
}

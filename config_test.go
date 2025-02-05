package gotel_test

import (
	"os"
	"testing"

	"github.com/lucavallin/gotel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfigFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		envVars map[string]string
		want    gotel.Config
		wantErr bool
	}{
		{
			name:    "success with default values",
			envVars: map[string]string{},
			want: gotel.Config{
				ServiceName:    "test-service",
				ServiceVersion: "0.0.1",
				Enabled:        true,
			},
			wantErr: false,
		},
		{
			name: "success with custom values",
			envVars: map[string]string{
				"SERVICE_NAME":      "test-service",
				"SERVICE_VERSION":   "1.0.0",
				"TELEMETRY_ENABLED": "true",
			},
			want: gotel.Config{
				ServiceName:    "test-service",
				ServiceVersion: "1.0.0",
				Enabled:        true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Clear environment before each test
			os.Clearenv()

			// Set environment variables for test
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}

			// Run test
			got, err := gotel.NewConfigFromEnv()
			if tt.wantErr {
				assert.Error(t, err)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/lucavallin/gotel"
)

func main() {
	ctx := context.Background()

	telemConfig, err := gotel.NewConfigFromEnv()
	if err != nil {
		fmt.Println("failed to load telemetry config")
		os.Exit(1)
	}

	// Initialize telemetry. If the exporter fails, fallback to nop.
	var telem gotel.TelemetryProvider
	telem, err = gotel.NewTelemetry(ctx, telemConfig)
	if err != nil {
		fmt.Println("failed to create telemetry, falling back to no-op telemetry")
		telem, _ = gotel.NewNoopTelemetry(telemConfig)
	}
	defer telem.Shutdown()

	telem.LogInfo("telemetry initialized")
}

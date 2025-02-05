# gotel

`gotel` is a Go library for handy observability. It is a wrapper around the OpenTelemetry Go library that provides a simplified API for instrumenting your Go applications.

`gotel` offers the following features:

- A simplified API for instrumenting your Go applications: log messages, metrics, and traces.
- OpenTelemetry exporters via OTLP gRPC
- A configuration struct that can be loaded from environment variables.
- A no-op telemetry provider that can be used as a fallback when the exporter fails.

## Installation

To install `gotel`, use `go get`:

```sh
go get github.com/lucavallin/gotel
```

## Usage

To use `gotel`, you need to create an instance `gotel.Telemetry` and use it to instrument your application. Here is an example:

```go
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
	defer telem.Shutdown(ctx)

	telem.LogInfo("telemetry initialized")
}
```

## Development

The repository contains a `justfile` with useful commands for development. To see the list of available commands, run `just`:

```sh
❯ just
Available recipes:
    build        # Build the gotel package
    default      # Default recipe
    deps         # Install dependencies
    gen-mocks    # Generate mocks
    help         # Show help message
    lint *ARGS   # Run linter (--fix to fix issues)
    test json="" # Run tests (--json to output coverage in json format)
```

## Contributing

Contributions are welcome! For bug reports, feature requests, or questions, please [open an issue](https://github.com/lucavallin/gotel/issues/new). For pull requests, fork the repository and submit a PR.

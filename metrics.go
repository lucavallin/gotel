package gotel

// Metric represents a metric that can be collected by the server.
type Metric struct {
	Name        string
	Unit        string
	Description string
}

// MetricRequestDurationMillis is a metric that measures the latency of HTTP requests processed by the server, in milliseconds.
var MetricRequestDurationMillis = Metric{
	Name:        "request_duration_millis",
	Unit:        "ms",
	Description: "Measures the latency of HTTP requests processed by the server, in milliseconds.",
}

// MetricRequestsInFlight is a metric that measures the number of requests currently being processed by the server.
var MetricRequestsInFlight = Metric{
	Name:        "requests_inflight",
	Unit:        "{count}",
	Description: "Measures the number of requests currently being processed by the server.",
}

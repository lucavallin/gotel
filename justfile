# Justfile
#
# To run a recipe:
#     just <recipe>
#
# Example:
#     just build
#

set unstable := true

# Default recipe
default: help

# Show help message
help:
    @just --list

# Build the gotel package
build:
    go build .

# Run linter (--fix to fix issues)
lint *ARGS:
	golangci-lint run {{ ARGS }}

# Run tests (--json to output coverage in json format)
[script]
test json="":
    if [ "{{ json }}" = "--json" ]; then
    	go test -cover -json -race -v ./... | tee coverage.json
    else
    	go test -race -v ./...
    fi

# Generate mocks
gen-mocks:
    mockery

# Install dependencies
[script]
deps:
	go install github.com/vektra/mockery/v2@v2.50.0
	go install github.com/mfridman/tparse@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)"/bin v1.61.0
	go mod download -x

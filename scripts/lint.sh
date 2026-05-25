#!/bin/bash
set -e
echo "=== Running Golang Lint ==="
golangci-lint run --timeout=5m
echo "=== Lint Completed ==="

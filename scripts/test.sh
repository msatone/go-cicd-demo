#!/bin/bash
set -e
echo "=== Running Unit Tests ==="
go test ./... -v -coverprofile=coverage.out
go tool cover -func=coverage.out
echo "=== Tests Completed ==="

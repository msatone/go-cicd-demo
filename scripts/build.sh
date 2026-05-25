#!/bin/bash
set -e
echo "=== Starting Go Build ==="
go mod tidy
go mod download
mkdir -p bin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o bin/my-go-app ./cmd/app
echo "=== Build Completed ==="
ls -lh bin/

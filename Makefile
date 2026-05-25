APP_NAME=my-go-app
BUILD_DIR=bin
DOCKER_IMAGE=my-go-app

.PHONY: fmt lint test build clean docker-build run

fmt:
	go fmt ./...

lint:
	golangci-lint run

test:
	go test ./... -v -coverprofile=coverage.out
	go tool cover -func=coverage.out

build:
	mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/app
	@echo "Binary built: $(BUILD_DIR)/$(APP_NAME)"

clean:
	rm -rf $(BUILD_DIR) coverage.out

docker-build:
	docker build -f deployment/Dockerfile -t $(DOCKER_IMAGE):local .

run:
	go run ./cmd/app/main.go

deps:
	go mod tidy
	go mod download

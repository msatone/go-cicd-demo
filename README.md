# my-go-app — Golang EKS Deployment via Harness

Golang HTTP application deployed to AWS EKS (Dummy-EKS) via Harness CI/CD.

## Project Structure

```
my-go-app/
├── cmd/app/main.go                  # Entry point, graceful shutdown
├── internal/
│   ├── handlers/health.go           # HTTP handlers for all endpoints
│   ├── services/health_service.go   # Business logic layer
│   └── models/response.go          # JSON response structs
├── pkg/utils/logger.go             # Zap logger
├── tests/integration_test.go       # Integration tests
├── deployment/
│   ├── Dockerfile                  # Multi-stage build
│   └── k8s/
│       ├── namespace.yaml          # my-app-ns namespace
│       ├── deployment.yaml         # 2 replicas, probes
│       └── service.yaml            # LoadBalancer port 80
├── .harness/
│   ├── ci-pipeline.yaml            # Build → Test → Lint → Push ECR
│   ├── cd-pipeline.yaml            # Deploy to Dummy-EKS
│   ├── service.yaml                # Harness service definition
│   └── environment.yaml            # Harness env + infra (Dummy-EKS)
├── scripts/
│   ├── build.sh
│   ├── test.sh
│   └── lint.sh
├── .golangci.yml                   # Lint rules
├── Makefile
└── go.mod
```

## Endpoints

| Endpoint  | Description                        |
|-----------|------------------------------------|
| `/`       | Home — app info, version, hostname |
| `/health` | Liveness probe                     |
| `/ready`  | Readiness probe                    |
| `/info`   | App name, Go version, hostname     |

## Local Run

```bash
# Install deps
go mod tidy

# Run tests
make test

# Run locally
make run

# Build binary
make build

# Build Docker image
make docker-build
```

## Harness Setup (one-time)

### Connectors needed
| Name                  | Type              | Purpose              |
|-----------------------|-------------------|----------------------|
| `harness_code_connector` | Harness Code   | Source code repo     |
| `aws_ecr_connector`   | AWS               | Push/pull ECR images |
| `eks_connector`       | Kubernetes        | Deploy to Dummy-EKS  |

### Replace in .harness/ files
- `YOUR_PROJECT_ID` → your Harness project identifier

## CI/CD Flow

```
Push to main
  → CI Pipeline triggers
  → Install deps → Format → Lint → Unit Tests → Security Scan → Build Binary
  → Docker Build → Push to ECR (tag: pipeline.sequenceId)
  → CD Pipeline triggers automatically
  → Apply namespace → Rolling Deploy to Dummy-EKS → Apply Service
  → Verify pods running → Print LoadBalancer URL
```

## Environment Variables

| Variable    | Default | Description       |
|-------------|---------|-------------------|
| PORT        | 8080    | Server listen port|
| APP_VERSION | 1.0.0   | Injected by Harness|

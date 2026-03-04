9jauni
======

Compact README for the 9jauni University API service.

Overview
-	This repository implements a small University API and a generated client. The project includes an OpenAPI spec, a Dockerfile and convenience compose manifests.

Features
-	HTTP API for university data (see `openapi.yaml` and `api/docs`)
-	A generated client in the `client/` folder
-	Dockerfile and `compose.yaml` for containerized runs

Prerequisites
-	Go (1.20+ recommended)
-	Docker (optional, for container builds)

Quick start (local)
- Build: `go build -o 9jauni ./`
- Run: `./9jauni`

Docker
- Build image: `docker build -t 9jauni:local .`
- Run container: `docker run --rm -p 8080:8080 9jauni:local`
- Or use compose: `docker compose up --build` (see `compose.yaml`)

OpenAPI & client
- API spec: `openapi.yaml`
- Generated client: see `client/` (includes example usage and `README.md`)

Tests
- Run unit tests: `go test ./...`
- API integration tests live in `test/`

Files of interest
- `main.go` — application entrypoint
- `openapi.yaml` — API contract
- `Dockerfile`, `compose.yaml` — containerization
- `client/` — generated client and helpers
- `json/uni.json` — sample data

Next steps
- Run `go test ./...` and try the Docker flow above
- Update the OpenAPI file and regenerate client if you change API

License
- None specified; add a LICENSE file if/when needed.

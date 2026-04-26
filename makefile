#!/bin/bash
# Run locally
run-local:
	go run cmd/server/main.go

# Build and run with Docker
build-docker:
	docker-compose up --build -d

# Run tests
test:
	go test ./...

# Stop and remove Docker containers
docker-down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f


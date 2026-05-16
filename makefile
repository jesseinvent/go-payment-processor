# Run locally
dev:
	air

# Run tests
test:
	go test ./...

# Install dependencies
tidy:
	go mod tidy

# Lint code
vet:
	go vet ./...

test:
	go test -v ./...	

# Format code
format:
	go fmt ./...

# Build binary
build:
	go build -o payment-processor main.go

# Run built binary
run:
	./payment-processor

## --- Docker commands ---

# Build and run with Docker
build-docker:
	docker compose up --build -d

# Stop and remove Docker containers
docker-down:
	docker compose down

# View logs
logs:
	docker compose logs -f

DB_URL=postgres://user:password@localhost:5432/payment_db?sslmode=disable

migrate-up:
	migrate -path ./migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path ./migrations -database "$(DB_URL)" down

migrate-status:
	migrate -path ./migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir ./migrations -seq $(name)
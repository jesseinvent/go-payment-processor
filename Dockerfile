FROM golang:latest

WORKDIR /app

COPY . .

ENV GIN_MODE=release
ENV ENVIRONMENT=production

# Install dependencies
RUN make tidy

# Install migrate binary
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Build the application
RUN make build

EXPOSE 5001

# Run migrations then start server
CMD ["sh", "-c", "migrate -path ./migrations -database $DB_URL up && make run"]
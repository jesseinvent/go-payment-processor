FROM golang:latest

WORKDIR /app

COPY . .

ENV GIN_MODE=release
ENV ENVIRONMENT=production

# Install dependencies
RUN make tidy

# Build the application
RUN make build

# Expose the port the app runs on
EXPOSE 5001

# Run the application
CMD ["make", "run"]
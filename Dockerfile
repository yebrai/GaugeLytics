# Build stage
FROM golang:1.22.5-alpine AS builder

WORKDIR /app

# Copy go.mod file
COPY go.mod ./
# Check if go.sum exists before copying it
RUN if [ -f go.sum ]; then cp go.sum .; fi

# Download dependencies
RUN go mod download

# Copy the source code and build the application
COPY . .
RUN go build -o gaugelytics ./cmd/gaugelytics/main.go

# Production stage
FROM alpine:3.18

WORKDIR /app
COPY --from=builder /app/gaugelytics .

# Expose the port for the app
EXPOSE 8080

# Command to run the app
CMD ["./gaugelytics"]

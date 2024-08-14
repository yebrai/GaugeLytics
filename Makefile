# Variables
APP_NAME = gaugelytics
DOCKER_COMPOSE_FILE = docker-compose.yml
ENV_FILE = .env

# Default target
.PHONY: all
all: build run

# Build the Go application
.PHONY: build
build:
	@echo "Building the Go application..."
	@go build -o ./cmd/$(APP_NAME)/$(APP_NAME) ./cmd/$(APP_NAME)/main.go

# Run the application using Docker Compose
.PHONY: run
run:
	@echo "Running the application using Docker Compose..."
	@docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) up --build

# Stop the application and remove containers
.PHONY: stop
stop:
	@echo "Stopping and removing containers..."
	@docker-compose --env-file $(ENV_FILE) -f $(DOCKER_COMPOSE_FILE) down

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test ./...

# Clean up build files
.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	@rm -f ./cmd/$(APP_NAME)/$(APP_NAME)

# Rebuild and rerun the application
.PHONY: rebuild
rebuild: clean build run

# Build Docker image
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME) .

# Push Docker image to a registry (Modify for your own registry)
.PHONY: docker-push
docker-push:
	@echo "Pushing Docker image to registry..."
	@docker tag $(APP_NAME):latest your-registry/$(APP_NAME):latest
	@docker push your-registry/$(APP_NAME):latest

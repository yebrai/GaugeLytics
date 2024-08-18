# Variables
SERVICES := health-service
DOCKER_COMPOSE_FILE := docker-compose.yml

# Default target
.PHONY: all
all: build test

# Build all services
.PHONY: build
build:
	@echo "Building all services..."
	@for service in $(SERVICES); do \
		docker-compose -f $(DOCKER_COMPOSE_FILE) build $$service ; \
	done

# Run tests for all services
.PHONY: test
test:
	@echo "Running tests for all services..."
	@go test ./test/...

# Run a specific service
.PHONY: run
run:
	@echo "Running all services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Stop all services
.PHONY: stop
stop:
	@echo "Stopping all services..."
	@docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Clean up build files
.PHONY: clean
clean:
	@echo "Cleaning up build files..."
	@for service in $(SERVICES); do \
		rm -f ./services/$$service/bin/* ; \
	done

# Deploy to production
.PHONY: deploy
deploy:
	@echo "Deploying to production..."

# Verify health-check endpoint
verify-health:
	@echo "Verifying health-check endpoint..."
	scripts/check_health.sh
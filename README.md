# GaugeLytics Backend

GaugeLytics is an application designed for tracking and analyzing web and mobile application analytics. This repository contains the backend services, built using Go and Docker.

# Technologies

- **Go:** The backend services are built with Go.
- **Docker:** Each service is containerized with Docker.
- **GitHub Actions:** CI/CD pipelines are automated with GitHub Actions for seamless integration, testing, and deployment.

## Project Structure

```plaintext
GaugeLytics-backend/
├── services/
│   ├── health-service/
│   │   ├── cmd/                      # Entry point for the service
│   │   ├── internal/
│   │   │   ├── application/          # Core application logic
│   │   │   ├── domain/               # Domain models and business rules
│   │   │   ├── infrastructure/       # Infrastructure-related code
│   │   │   └── config/               # Configuration management
│   │   ├── test/health-service/      # Tests for health service
│   │   │   ├── integration/          # Integration tests
│   │   │   └── unit/                 # Unit tests
│   │   ├── Dockerfile                # Dockerfile for the health service
│   │   ├── go.mod                    # Go module file
│   │   └── go.sum                    # Go checksum file
│   └── user-service/
│       ├── cmd/                      # Entry point for the user service
│       ├── internal/
│       │   ├── application/          # Core application logic
│       │   ├── domain/               # Domain models and business rules
│       │   ├── infrastructure/       # Infrastructure-related code
│       │   ├── ports/                # Ports and adapters
│       ├── test/user-service/        # Tests for user service
│       │   ├── integration/          # Integration tests
│       │   └── unit/                 # Unit tests
│       ├── Dockerfile                # Dockerfile for the user service
│       ├── go.mod                    # Go module file
│       └── go.sum                    # Go checksum file
├── scripts/                          # Utility scripts for the project
├── Dockerfile                        # Dockerfile for the overall application setup
├── Makefile                          # Makefile with tasks for building, running, and testing services
├── .gitignore                        # Git ignore file
├── README.md                         # Project documentation
└── docker-compose.yml                # Docker Compose file for managing multi-container deployment
```
### Getting Started

Building the Project
To build all the services:

```
make build
```
### Running the Project
To start all the services:

```
make run
```
### Testing the Services
To run tests:

```
make test
```
### Stopping the Services
To stop all running services:

```
make stop
```
### Cleaning Up Build Files
To clean up build artifacts:

```
make clean
```
Environment Variables
Use a .env file in the root directory for environment-specific configurations.

### License
This project is licensed under the MIT License.

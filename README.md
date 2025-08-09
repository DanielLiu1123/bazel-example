# Bazel Multi-Language Monorepo Example

This is a multi-language monorepo built with Bazel, supporting Java and Go services with Protocol Buffers for inter-service communication.

## Project Structure

```
bazel-example/
├── MODULE.bazel              # Bazel module configuration (Bzlmod)
├── BUILD.bazel               # Root build file
├── .bazelrc                  # Bazel configuration
├── go.mod                    # Go module definition
│
├── proto/                    # Protocol Buffer definitions
│   ├── common/               # Shared message types
│   ├── user/                 # User service API
│   ├── order/                # Order service API
│   ├── gateway/              # Gateway API
│   ├── notification/         # Notification service API
│   ├── worker/               # Worker service API
│   └── internal/             # Internal service communication
│
├── services/                 # Microservices (organized by business function)
│   ├── user/                 # User management service (Java)
│   │   ├── src/              # Java source code
│   │   ├── config/           # Service configuration
│   │   └── BUILD.bazel       # Build configuration
│   ├── order/                # Order management service (Java)
│   │   ├── src/              # Java source code
│   │   ├── config/           # Service configuration
│   │   └── BUILD.bazel       # Build configuration
│   ├── gateway/              # API Gateway (Go)
│   │   ├── main.go           # Main application
│   │   ├── internal/         # Internal packages
│   │   ├── config/           # Service configuration
│   │   └── BUILD.bazel       # Build configuration
│   ├── notification/         # Notification service (Go)
│   └── worker/               # Background job processor (Go)
│
├── libs/                     # Shared libraries
│   ├── common/               # Common utilities
│   ├── database/             # Database utilities
│   ├── messaging/            # Message queue utilities
│   └── monitoring/           # Monitoring and metrics
│
├── tools/                    # Build and development tools
├── deployments/              # Deployment configurations
└── docs/                     # Documentation
```

## Features

- **Multi-language support**: Java and Go services in the same repository
- **Protocol Buffers**: Type-safe inter-service communication
- **Bazel build system**: Fast, reliable, and scalable builds
- **Microservices architecture**: Services organized by business function
- **Shared libraries**: Reusable code across services
- **Docker support**: Containerized deployment
- **Kubernetes ready**: K8s deployment configurations

## Getting Started

### Prerequisites

- Bazel 7.0+ (see `.bazelversion`)
- Java 17+
- Go 1.21+
- Docker (optional, for containerized deployment)

### Building the Project

```bash
# Build all services
./tools/scripts/build.sh

# Or build specific services
bazel build //services/user/java:user_service
bazel build //services/gateway/go:gateway
```

### Running Tests

```bash
# Run all tests
./tools/scripts/test.sh

# Or run specific tests
bazel test //services/user/java:user_service_test
bazel test //services/gateway/go:gateway_test
```

### Local Development with Docker Compose

```bash
cd deployments/docker-compose
docker-compose up -d
```

## Services

### User Service (Java)
- **Port**: 8081
- **Technology**: Spring Boot + gRPC
- **Purpose**: User management and authentication

### Order Service (Java)
- **Port**: 8082
- **Technology**: Spring Boot + gRPC
- **Purpose**: Order processing and management

### Gateway (Go)
- **Port**: 8080
- **Technology**: Gin + gRPC clients
- **Purpose**: API gateway and request routing

### Notification Service (Go)
- **Technology**: gRPC
- **Purpose**: Send notifications to users

### Worker Service (Go)
- **Technology**: gRPC
- **Purpose**: Background job processing

## Development

### Adding a New Service

1. Create service directory under `services/`
2. Add proto definitions under `proto/`
3. Create BUILD.bazel files for build configuration
4. Implement service logic in chosen language
5. Update deployment configurations

### Adding Dependencies

- **Java**: Update `MODULE.bazel` maven.install section
- **Go**: Update `go.mod` and run `bazel run //:gazelle-update-repos`

## Architecture

This monorepo follows a microservices architecture with:

- **API Gateway**: Single entry point for all client requests
- **Business Services**: Domain-specific services (user, order)
- **Support Services**: Cross-cutting concerns (notification, worker)
- **Shared Libraries**: Common utilities and protocols

## Contributing

1. Follow the existing directory structure
2. Add appropriate BUILD.bazel files for new code
3. Update proto definitions for API changes
4. Add tests for new functionality
5. Update documentation as needed

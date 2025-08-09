# Getting Started with Development

This guide will help you set up the development environment for the Bazel multi-language monorepo.

## Prerequisites

### Required Tools

1. **Bazel 7.0+**
   ```bash
   # macOS
   brew install bazel
   
   # Linux
   curl -L https://github.com/bazelbuild/bazel/releases/download/7.0.0/bazel-7.0.0-linux-x86_64 -o /usr/local/bin/bazel
   chmod +x /usr/local/bin/bazel
   ```

2. **Java 17+**
   ```bash
   # macOS
   brew install openjdk@17
   
   # Linux
   sudo apt-get install openjdk-17-jdk
   ```

3. **Go 1.21+**
   ```bash
   # macOS
   brew install go
   
   # Linux
   wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
   sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
   ```

4. **Docker** (optional, for containerized development)
   ```bash
   # macOS
   brew install docker
   
   # Linux
   sudo apt-get install docker.io
   ```

### Optional Tools

- **PostgreSQL** (for local database development)
- **Redis** (for caching)
- **kubectl** (for Kubernetes deployment)

## Project Setup

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd bazel-example
   ```

2. **Initialize Go modules**
   ```bash
   go mod tidy
   ```

3. **Generate Bazel files for Go**
   ```bash
   bazel run //:gazelle
   ```

## Building the Project

### Build All Services

```bash
# Using the build script
./tools/scripts/build.sh

# Or manually
bazel build //services/user:user_service
bazel build //services/order:order_service
bazel build //services/gateway:gateway
bazel build //services/notification:notification
bazel build //services/worker:worker
```

### Build Specific Services

```bash
# Java services
bazel build //services/user:user_service
bazel build //services/order:order_service

# Go services
bazel build //services/gateway:gateway
bazel build //services/notification:notification
bazel build //services/worker:worker
```

### Build Protocol Buffers

```bash
# Build all proto libraries
bazel build //proto/...

# Build specific proto libraries
bazel build //proto/user:user_java_proto
bazel build //proto/common:types_go_proto
```

## Running Tests

### Run All Tests

```bash
# Using the test script
./tools/scripts/test.sh

# Or manually
bazel test //...
```

### Run Specific Tests

```bash
# Java service tests
bazel test //services/user:user_service_test
bazel test //services/order:order_service_test

# Go service tests
bazel test //services/gateway:gateway_test
bazel test //services/notification:notification_test
```

## Local Development

### Using Docker Compose

1. **Start infrastructure services**
   ```bash
   cd deployments/docker-compose
   docker-compose up -d postgres redis
   ```

2. **Build and run services**
   ```bash
   # Build services first
   ./tools/scripts/build.sh
   
   # Run services
   docker-compose up
   ```

### Running Services Individually

1. **Start infrastructure**
   ```bash
   # PostgreSQL
   docker run -d --name postgres \
     -e POSTGRES_DB=bazel_example \
     -e POSTGRES_USER=postgres \
     -e POSTGRES_PASSWORD=password \
     -p 5432:5432 postgres:15
   
   # Redis
   docker run -d --name redis -p 6379:6379 redis:7-alpine
   ```

2. **Run Java services**
   ```bash
   # User service
   bazel run //services/user:user_service
   
   # Order service
   bazel run //services/order:order_service
   ```

3. **Run Go services**
   ```bash
   # Gateway
   bazel run //services/gateway:gateway
   
   # Notification service
   bazel run //services/notification:notification
   
   # Worker service
   bazel run //services/worker:worker
   ```

## Development Workflow

### Adding New Dependencies

#### Java Dependencies
1. Update `MODULE.bazel` in the `maven.install` section
2. Rebuild the project

#### Go Dependencies
1. Add to `go.mod`
2. Run `go mod tidy`
3. Update Bazel files: `bazel run //:gazelle-update-repos`

### Adding New Services

1. **Create service directory**
   ```bash
   mkdir -p services/new-service/src/main/java/com/example/newservice
   ```

2. **Add proto definitions**
   ```bash
   mkdir -p proto/new-service
   # Add .proto files
   ```

3. **Create BUILD.bazel files**
   - Add build configurations for the new service
   - Update proto BUILD files

4. **Update build scripts**
   - Add new service to `tools/scripts/build.sh`
   - Add tests to `tools/scripts/test.sh`

### Code Style and Formatting

- **Java**: Follow Google Java Style Guide
- **Go**: Use `gofmt` and `golint`
- **Proto**: Follow Protocol Buffers Style Guide

### Debugging

#### Java Services
```bash
# Run with debug port
bazel run //services/user:user_service -- --debug
```

#### Go Services
```bash
# Run with delve debugger
dlv exec bazel-bin/services/gateway/gateway_/gateway
```

## Troubleshooting

### Common Issues

1. **Bazel build failures**
   - Clean build cache: `bazel clean`
   - Check Java/Go versions
   - Verify dependencies in MODULE.bazel

2. **Proto generation issues**
   - Regenerate: `bazel build //proto/...`
   - Check proto syntax

3. **Database connection issues**
   - Verify PostgreSQL is running
   - Check connection strings in configuration

### Getting Help

- Check the [Architecture Documentation](../architecture/overview.md)
- Review [API Documentation](../api/README.md)
- Open an issue in the repository

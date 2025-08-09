# Multi-stage Dockerfile for Go services
FROM golang:1.21-alpine as builder

# Install Bazel
RUN apk add --no-cache bash curl && \
    curl -L https://github.com/bazelbuild/bazel/releases/download/7.0.0/bazel-7.0.0-linux-x86_64 -o /usr/local/bin/bazel && \
    chmod +x /usr/local/bin/bazel

# Set working directory
WORKDIR /workspace

# Copy source code
COPY . .

# Build argument for service target
ARG SERVICE_TARGET

# Build the service
RUN bazel build ${SERVICE_TARGET}

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Create app user
RUN addgroup -g 1001 app && adduser -D -s /bin/sh -u 1001 -G app app

# Set working directory
WORKDIR /app

# Copy built binary from builder stage
ARG SERVICE_TARGET
COPY --from=builder /workspace/bazel-bin/${SERVICE_TARGET#//} ./app

# Change ownership
RUN chown -R app:app /app

# Switch to app user
USER app

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the application
ENTRYPOINT ["./app"]

# Multi-stage Dockerfile for Java services
FROM openjdk:17-jdk-slim as builder

# Install Bazel
RUN apt-get update && apt-get install -y curl gnupg && \
    curl -fsSL https://bazel.build/bazel-release.pub.gpg | gpg --dearmor > bazel.gpg && \
    mv bazel.gpg /etc/apt/trusted.gpg.d/ && \
    echo "deb [arch=amd64] https://storage.googleapis.com/bazel-apt stable jdk1.8" | tee /etc/apt/sources.list.d/bazel.list && \
    apt-get update && apt-get install -y bazel

# Set working directory
WORKDIR /workspace

# Copy source code
COPY . .

# Build argument for service target
ARG SERVICE_TARGET

# Build the service
RUN bazel build ${SERVICE_TARGET}

# Runtime stage
FROM openjdk:17-jre-slim

# Create app user
RUN groupadd -r app && useradd -r -g app app

# Set working directory
WORKDIR /app

# Copy built JAR from builder stage
ARG SERVICE_TARGET
COPY --from=builder /workspace/bazel-bin/${SERVICE_TARGET#//} ./app.jar

# Change ownership
RUN chown -R app:app /app

# Switch to app user
USER app

# Expose port
EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/actuator/health || exit 1

# Run the application
ENTRYPOINT ["java", "-jar", "app.jar"]

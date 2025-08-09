#!/bin/bash

# Build script for the entire project
set -e

echo "Building all services..."

# Build Java services
echo "Building Java services..."
bazel build //services/user:user_service
bazel build //services/order:order_service

# Build Go services
echo "Building Go services..."
bazel build //services/gateway:gateway
bazel build //services/notification:notification
bazel build //services/worker:worker

echo "Build completed successfully!"

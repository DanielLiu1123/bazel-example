#!/bin/bash

# Test script for the entire project
set -e

echo "Running all tests..."

# Test Java services
echo "Testing Java services..."
bazel test //services/user:user_service_test
bazel test //services/order:order_service_test

# Test Go services
echo "Testing Go services..."
bazel test //services/gateway:gateway_test
bazel test //services/notification:notification_test
bazel test //services/worker:worker_test

echo "All tests passed!"

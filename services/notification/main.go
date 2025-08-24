package main

import "github.com/example/bazel-example/services/notification/internal/grpc"

func main() {
	ns := grpc.NewNotificationService()
	println(ns)
	println("Hello, World!")
}

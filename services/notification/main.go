package main

import "github.com/example/bazel-example/services/notification/grpc"

func main() {
	ns := grpc.NewNotificationService()
	println(ns)
	println("Hello, World!")
}

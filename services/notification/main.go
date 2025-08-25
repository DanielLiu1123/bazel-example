package main

import (
	"fmt"
	"net"

	"github.com/example/bazel-example/proto/notificationpb"
	grpc2 "github.com/example/bazel-example/services/notification/grpc"
	"google.golang.org/grpc"
)

func main() {

	// Create a new gRPC server
	server := grpc.NewServer()

	notificationpb.RegisterNotificationServiceServer(server, grpc2.NewNotificationService())

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server started on port 9090")

	err = server.Serve(lis)
	if err != nil {
		panic(err)
	}
}

package grpc

import (
	notificationgrpc "github.com/example/bazel-example/proto/notificationgrpc"
)

type NotificationService struct {
	notificationgrpc.UnimplementedNotificationServiceServer
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

package grpc

import (
	"github.com/example/bazel-example/proto/notificationpb"
)

type NotificationService struct {
	notificationpb.UnimplementedNotificationServiceServer
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

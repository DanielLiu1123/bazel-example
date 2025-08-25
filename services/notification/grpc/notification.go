package grpc

import (
	"context"

	"github.com/example/bazel-example/proto/notificationpb"
)

type NotificationService struct {
	notificationpb.UnimplementedNotificationServiceServer
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n NotificationService) CreateNotification(ctx context.Context, request *notificationpb.CreateNotificationRequest) (*notificationpb.CreateNotificationResponse, error) {
	//TODO implement me
	panic("implement me")
}

package client

import (
	"context"
	"time"

	"github.com/htchoi-nurilab/go-msa/notification-service/proto"
)

type NotificationGrpcClient struct {
	c       notificationpb.NotificationServiceClient
	timeout time.Duration
}

func NewGrpcNotificationClient(c notificationpb.NotificationServiceClient, timeout time.Duration) *NotificationGrpcClient {
	return &NotificationGrpcClient{c: c, timeout: timeout}
}

func (g *NotificationGrpcClient) CreateWelcomeNotification(ctx context.Context, userID uint, name string) error {
	grpcCtx, cancel := context.WithTimeout(ctx, g.timeout)
	defer cancel()

	_, err := g.c.CreateNotification(grpcCtx, &notificationpb.NotificationCreateRequest{
		UserId: int64(userID),
		Name:   name,
	})

	return err
}

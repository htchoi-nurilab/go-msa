package client

import "context"

type NotificationClient interface {
	CreateWelcomeNotification(ctx context.Context, userID uint, name string) error
}

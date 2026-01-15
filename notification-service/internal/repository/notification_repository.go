package repository

import (
	"context"

	"github.com/htchoi-nurilab/go-msa/notification-service/internal/domain"
)

type NotificationRepository interface {
	Save(ctx context.Context, notification *domain.Notification) error
}

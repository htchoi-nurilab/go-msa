package repository

import (
	"context"

	"github.com/htchoi-nurilab/go-msa/notification-service/internal/domain"
	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) Save(ctx context.Context, notification *domain.Notification) error {
	return r.db.WithContext(ctx).Create(&notification).Error
}

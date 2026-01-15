package service

import (
	"context"
	"fmt"

	"github.com/htchoi-nurilab/go-msa/notification-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/notification-service/internal/repository"
)

type NotificationService struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationService(notificationRepo repository.NotificationRepository) *NotificationService {
	return &NotificationService{notificationRepo: notificationRepo}
}

func (s *NotificationService) CreateWelcomeNotification(ctx context.Context, userID uint, name string) error {
	message := fmt.Sprintf("%s님 환영합니다", name)

	n := &domain.Notification{
		UserID:  userID,
		Message: message,
	}

	return s.notificationRepo.Save(ctx, n)
}

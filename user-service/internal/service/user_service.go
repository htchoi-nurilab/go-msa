package service

import (
	"context"
	"log"

	"github.com/htchoi-nurilab/go-msa/user-service/internal/client"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/dto"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo   repository.UserRepository
	notiClient client.NotificationClient
}

func NewUserService(userRepo repository.UserRepository, notiClient client.NotificationClient) *UserService {
	return &UserService{userRepo: userRepo, notiClient: notiClient}
}

func (s *UserService) CreateUser(ctx context.Context, request dto.UserCreateRequest) (dto.UserCreateResponse, error) {
	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return dto.UserCreateResponse{}, err
	}

	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedPassword,
	}

	if err := s.userRepo.Save(&user); err != nil {
		return dto.UserCreateResponse{}, err
	}

	if s.notiClient != nil {
		if err := s.notiClient.CreateWelcomeNotification(ctx, user.ID, user.Name); err != nil {
			log.Println("failed to create notification:", err)
		}
	}

	return dto.UserCreateResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

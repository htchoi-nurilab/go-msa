package service

import (
	"github.com/htchoi-nurilab/go-msa/user-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/dto"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (UserSvc *UserService) CreateUser(request dto.UserCreateRequest) (dto.UserCreateResponse, error) {
	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
	}

	if err := UserSvc.userRepo.Save(&user); err != nil {
		return dto.UserCreateResponse{}, err
	}

	return dto.UserCreateResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}

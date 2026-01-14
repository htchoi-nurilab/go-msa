package service

import (
	"github.com/htchoi-nurilab/go-msa/user-service/internal/domain"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/dto"
	"github.com/htchoi-nurilab/go-msa/user-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (UserSvc *UserService) CreateUser(request dto.UserCreateRequest) (dto.UserCreateResponse, error) {
	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return dto.UserCreateResponse{}, err
	}

	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: hashedPassword,
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

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

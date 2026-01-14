package repository

import "github.com/htchoi-nurilab/go-msa/user-service/internal/domain"

type UserRepository interface {
	Save(user *domain.User) error
}

package service

import (
	"github.com/vtdthang/goapi/entities"
	"github.com/vtdthang/goapi/user/repository"
)

// IUserService represent the user usecases
type IUserService interface {
	FindByEmail(email string) (*entities.User, error)
}

type userService struct {
	userRepo repository.IUserRepository
}

// NewUserService will create new an userService object representation of IUserService interface
func NewUserService(userRepo repository.IUserRepository) IUserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) FindByEmail(email string) (*entities.User, error) {
	res, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return res, nil
}

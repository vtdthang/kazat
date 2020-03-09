package service

import (
	"fmt"
	"net/http"

	"github.com/rs/xid"
	"github.com/vtdthang/goapi/entities"
	"github.com/vtdthang/goapi/lib/auth"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
	"github.com/vtdthang/goapi/models"
	"github.com/vtdthang/goapi/user/repository"
)

// IUserService represent the user usecases
type IUserService interface {
	FindByEmail(email string) (*entities.User, error)
	Register(registerRequest models.UserRegisterRequest) (*entities.User, error)
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

func (u *userService) Register(registerRequest models.UserRegisterRequest) (*entities.User, error) {
	existingUser, err := u.userRepo.FindByEmail(registerRequest.Email)

	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		err := httperror.NewHTTPError(http.StatusBadRequest, enums.EmailAlreadyExistsErrCode, enums.EmailAlreadyExistsErrMsg)

		return nil, err
	}

	userID := xid.New().String()
	accessToken, err := auth.GenerateJwtToken(userID)
	if err != nil {
		err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)

		return nil, err
	}

	newUser := &entities.User{
		ID:        userID,
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Email:     registerRequest.Email,
	}

	fmt.Println(newUser.ID)
	fmt.Println(accessToken)

	return newUser, nil
}

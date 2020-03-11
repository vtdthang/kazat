package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/vtdthang/goapi/entities"
	"github.com/vtdthang/goapi/lib/constants"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
	"github.com/vtdthang/goapi/lib/helpers"
	"github.com/vtdthang/goapi/models"
	"github.com/vtdthang/goapi/user/repository"
)

// IUserService represent the user usecases
type IUserService interface {
	FindByEmail(email string) (*entities.User, error)
	Login(userRequest models.UserLoginRequest) (*models.UserLoginOrRegisterResponse, error)
	Register(registerRequest models.UserRegisterRequest) (*models.UserLoginOrRegisterResponse, error)
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

func (u *userService) Login(userLoginRequest models.UserLoginRequest) (*models.UserLoginOrRegisterResponse, error) {
	existingUserEntity, err := u.userRepo.FindByEmail(userLoginRequest.Username)
	if err != nil {
		return nil, err
	}

	if existingUserEntity == nil {
		err := httperror.NewHTTPError(http.StatusBadRequest, enums.UserNotFoundErrCode, enums.UserNotFoundErrMsg)
		return nil, err
	}

	isPasswordMatched := helpers.ComparePassword(existingUserEntity.Password, []byte(userLoginRequest.Password))
	if !isPasswordMatched {
		err := httperror.NewHTTPError(http.StatusBadRequest, enums.PasswordIsIncorrectErrCode, enums.PasswordIsIncorrectErrMsg)
		return nil, err
	}

	accessToken, err := helpers.GenerateJwtToken(existingUserEntity.ID)
	if err != nil {
		err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
		return nil, err
	}

	refresToken, err := helpers.GenerateRandomString(32)
	if err != nil {
		fmt.Println("Generate ramdom byte is fail")
		return nil, err
	}

	newAuthenticationData := entities.Auth{
		ID:           models.NewID(),
		UserID:       existingUserEntity.ID,
		RefreshToken: refresToken,
		ExpiresAt:    models.GetMillisecondsForSpecificTime(time.Now().AddDate(0, 0, 60)),
		CreatedAt:    models.GetMilliseconds(),
		UpdatedAt:    models.GetMilliseconds(),
	}

	err = u.userRepo.InsertOneAuthData(newAuthenticationData)
	if err != nil {
		return nil, err
	}

	loginUserResponse := &models.UserLoginOrRegisterResponse{
		AccessToken:  accessToken,
		TokenType:    constants.SystemJWTTokenType,
		ExpiresIn:    constants.SystemJWTExpiresIn,
		RefreshToken: refresToken,
		UserProfile: models.UserProfileResponse{
			ID:        existingUserEntity.ID,
			FirstName: existingUserEntity.FirstName,
			LastName:  existingUserEntity.LastName,
			Email:     existingUserEntity.Email,
		},
	}

	return loginUserResponse, nil
}

func (u *userService) Register(registerRequest models.UserRegisterRequest) (*models.UserLoginOrRegisterResponse, error) {
	existingUser, err := u.userRepo.FindByEmail(strings.ToLower(registerRequest.Email))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if existingUser != nil {
		err := httperror.NewHTTPError(http.StatusBadRequest, enums.EmailAlreadyExistsErrCode, enums.EmailAlreadyExistsErrMsg)
		fmt.Println(err)
		return nil, err
	}

	userID := models.NewID()
	accessToken, err := helpers.GenerateJwtToken(userID)
	if err != nil {
		err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
		return nil, err
	}

	newUser := entities.User{
		ID:        userID,
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Email:     strings.ToLower(registerRequest.Email),
		CreatedAt: models.GetMilliseconds(),
		UpdatedAt: models.GetMilliseconds(),
	}

	hashedPassword, err := helpers.HashPassword([]byte(registerRequest.Password))
	if err != nil {
		err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
		return nil, err
	}

	newUser.Password = hashedPassword

	refresToken, err := helpers.GenerateRandomString(32)
	if err != nil {
		fmt.Println("Generate refresh token failed")
	}

	newAuthenticationData := entities.Auth{
		ID:           models.NewID(),
		UserID:       userID,
		RefreshToken: refresToken,
		ExpiresAt:    models.GetMillisecondsForSpecificTime(time.Now().AddDate(0, 0, 60)),
		CreatedAt:    models.GetMilliseconds(),
		UpdatedAt:    models.GetMilliseconds(),
	}

	err = u.userRepo.CreateUserAndAuthData(newUser, newAuthenticationData)
	if err != nil {
		return nil, err
	}

	registerUserResponse := &models.UserLoginOrRegisterResponse{
		AccessToken:  accessToken,
		TokenType:    constants.SystemJWTTokenType,
		ExpiresIn:    constants.SystemJWTExpiresIn,
		RefreshToken: refresToken,
		UserProfile: models.UserProfileResponse{
			ID:        userID,
			FirstName: registerRequest.FirstName,
			LastName:  registerRequest.LastName,
			Email:     registerRequest.Email,
		},
	}

	return registerUserResponse, nil
}

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vtdthang/goapi/lib/auth"
	apiresponse "github.com/vtdthang/goapi/lib/infrastructure"
	"github.com/vtdthang/goapi/models"
	"github.com/vtdthang/goapi/user/service"
)

// UserController  represent the http handler for user
type UserController struct {
	UserService service.IUserService
}

// NewUserController will initialize the users/ resources endpoint
func NewUserController(u service.IUserService) *UserController {
	return &UserController{UserService: u}
}

// FindByEmail find user by email
func (u *UserController) FindByEmail(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	email := "test1@gmail.com"

	user, err := u.UserService.FindByEmail(email)

	if err != nil {
		apiresponse.AsErrorResponse(w, err)
	} else {
		apiresponse.AsSuccessResponse(w, user)
	}
}

// Register to use create new account
func (u *UserController) Register(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)

	var registerReqModel models.UserRegisterRequest

	err := decoder.Decode(&registerReqModel)

	if err != nil {
		fmt.Println(err)
	}

	userRegisterResponse, err := u.UserService.Register(registerReqModel)

	fmt.Println(userRegisterResponse)

	userID := "1"

	user, err := u.UserService.FindByEmail(registerReqModel.Email)
	if err != nil {
		token, err := auth.GenerateJwtToken(userID)
		if err != nil {
			fmt.Println("ERROR GENERATE TOKEN")
		}

		fmt.Println(token)
	}

	if err != nil {
		apiresponse.AsErrorResponse(w, err)
	} else {
		apiresponse.AsSuccessResponse(w, user)
	}
}

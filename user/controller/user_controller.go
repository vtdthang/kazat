package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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

// Register to use create new account
func (u *UserController) Register(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)

	var registerReqModel models.UserRegisterRequest

	err := decoder.Decode(&registerReqModel)

	if err != nil {
		fmt.Println(err)
	}

	userRegisterResponse, err := u.UserService.Register(registerReqModel)

	if err != nil {
		apiresponse.AsErrorResponse(w, err)
	} else {
		apiresponse.AsSuccessResponse(w, userRegisterResponse)
	}
}

// Login api
func (u *UserController) Login(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(req.Body)

	var loginReqModel models.UserLoginRequest

	err := decoder.Decode(&loginReqModel)

	if err != nil {
		fmt.Println(err)
		apiresponse.AsErrorResponse(w, err)
	}

	userLoginResponse, err := u.UserService.Register(registerReqModel)

	if err != nil {
		apiresponse.AsErrorResponse(w, err)
	} else {
		apiresponse.AsSuccessResponse(w, userLoginResponse)
	}
}

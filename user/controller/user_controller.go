package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	email := "test2@gmail.com"

	user, err := u.UserService.FindByEmail(email)

	if err != nil {
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message": "Failed"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(user)
	}
}

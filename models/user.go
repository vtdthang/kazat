package models

// UserRegisterRequest model
type UserRegisterRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

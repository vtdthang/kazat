package models

// UserRegisterRequest model
type UserRegisterRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

// UserLoginRequest request
type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserProfileResponse user profile response
type UserProfileResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// UserLoginOrRegisterResponse user register response
type UserLoginOrRegisterResponse struct {
	AccessToken  string              `json:"access_token"`
	RefreshToken string              `json:"refresh_token"`
	TokenType    string              `json:"token_type"`
	ExpiresIn    int                 `json:"expires_in"`
	UserProfile  UserProfileResponse `json:"user_profile"`
}

package entities

// Auth is table to define user authentication
type Auth struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

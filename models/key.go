package models

// ContextKey contains request scoped variable
type ContextKey string

//
const (
	ContextKeyUserID  = ContextKey("user_id")
	ContextKeyAnother = ContextKey("another")
)

//
const (
	EnvJWTSecretKey = "JWT_SECRET_KEY"
	JWTSubjectKey   = "sub"
	JWTIssuedAtKey  = "iat"
	JWTExpiresAtKey = "exp"
	JWTIssuerKey    = "iss"
)

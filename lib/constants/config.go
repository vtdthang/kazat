package constants

//
const (
	SystemJWTTokenType = "Bearer"
	SystemJWTExpiresIn = 7200 // seconds

	EnvJWTSecretKey = "JWT_SECRET_KEY"
	EnvPostgresURL  = "PG_URL"
)

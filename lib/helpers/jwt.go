package helpers

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vtdthang/goapi/models"
)

// GenerateJwtToken generate signed encrypted jwt token
func GenerateJwtToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims[models.JWTSubjectKey] = userID // this token belongs to whom
	claims[models.JWTIssuedAtKey] = time.Now().Unix()
	claims[models.JWTExpiresAtKey] = time.Now().Add(time.Minute * 1).Unix() //Token expires after 1 hour
	claims[models.JWTIssuerKey] = "developer.kazat.com"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv(models.EnvJWTSecretKey)))
}

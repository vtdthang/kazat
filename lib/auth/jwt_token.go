package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/vtdthang/goapi/lib/helpers"
)

const jwtSecretKey = "ABCD123@"

// GenerateJwtToken generate signed encrypted jwt token
func GenerateJwtToken(userID string) (string, error) {
	defer helpers.TimeTrack(time.Now(), "GenerateJwtToken")

	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix() //Token expires after 1 hour
	claims["iss"] = "developer.kazat.com"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecretKey))
}

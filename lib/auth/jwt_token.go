package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const jwtSecretKey = "ABCD123@"

// GenerateJwtToken generate signed encrypted jwt token
func GenerateJwtToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecretKey))
}
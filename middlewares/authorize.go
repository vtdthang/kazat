package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
	apiresponse "github.com/vtdthang/goapi/lib/infrastructure"
	"github.com/vtdthang/goapi/models"
)

const authorizationHeader = "Authorization"

type decodedJWTToken struct {
	DeviceID string `json:"device_id"`
	jwt.StandardClaims
}

// AuthorizeMiddleware secure api endpoints
func AuthorizeMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authorizationHeader := req.Header.Get(authorizationHeader)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				decodedToken := &decodedJWTToken{}
				token, err := parseTokenWithClaim(bearerToken[1], decodedToken)

				if err != nil {
					v, _ := err.(*jwt.ValidationError)
					if v.Errors == jwt.ValidationErrorExpired {
						err := httperror.NewHTTPError(http.StatusUnauthorized, enums.TokenIsExpiredErrCode, enums.TokenIsExpiredErrMsg)
						apiresponse.AsErrorResponse(w, err)
						return
					}

					err := httperror.NewHTTPError(http.StatusForbidden, enums.TokenIsMalformedOrInvalidErrCode, enums.TokenIsMalformedOrInvalidErrMsg)
					apiresponse.AsErrorResponse(w, err)
					return
				}

				if !token.Valid {
					err := httperror.NewHTTPError(http.StatusForbidden, enums.TokenIsMalformedOrInvalidErrCode, enums.TokenIsMalformedOrInvalidErrMsg)
					apiresponse.AsErrorResponse(w, err)
					return
				}

				ctx := context.WithValue(req.Context(), models.ContextKeyUserID, decodedToken.Subject)
				req = req.WithContext(ctx)

				next(w, req, ps)
				return
			}

			err := httperror.NewHTTPError(http.StatusForbidden, enums.AuthorizationHeaderIsRequiredErrCode, enums.AuthorizationHeaderIsRequiredErrCMsg)
			apiresponse.AsErrorResponse(w, err)
			return
		}

		err := httperror.NewHTTPError(http.StatusForbidden, enums.AuthorizationHeaderIsRequiredErrCode, enums.AuthorizationHeaderIsRequiredErrCMsg)
		apiresponse.AsErrorResponse(w, err)
		return
	}
}

func parseBearerToken(bearerToken string) (*jwt.Token, error) {
	return jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte(os.Getenv(models.EnvJWTSecretKey)), nil
	})
}

func parseTokenWithClaim(bearerToken string, claims *decodedJWTToken) (*jwt.Token, error) {
	return jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv(models.EnvJWTSecretKey)), nil
	})
}

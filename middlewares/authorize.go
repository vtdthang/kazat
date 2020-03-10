package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
	apiresponse "github.com/vtdthang/goapi/lib/infrastructure"
)

const secretKey = "ABCD123@"
const authorizationHeader = "Authorization"

// AuthorizeMiddleware secure api endpoints
func AuthorizeMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authorizationHeader := req.Header.Get(authorizationHeader)
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := parseBearerToken(bearerToken[1])

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

				if token.Valid {
					//context.Set(req, "decoded", token.Claims)

					claims := token.Claims
					fmt.Println("CLAIMS ", claims)
					next(w, req, ps)
					return
				}
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
		return []byte(secretKey), nil
	})
}

// Exception blah blah blah
type Exception struct {
	Message string `json:"message"`
}

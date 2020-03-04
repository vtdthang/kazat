package httperror

import (
	"strconv"

	"github.com/vtdthang/goapi/lib/enums"
)

// HTTPError define all api structured errors
type HTTPError struct {
	StatusCode   int                `json:"-"`
	ErrorCode    enums.ErrorCode    `json:"error_code"`
	ErrorMessage enums.ErrorMessage `json:"error_message"`
}

//
func (e *HTTPError) Error() string {
	return strconv.Itoa(e.StatusCode)
}

// NewHTTPError create an instance of HttpError
func NewHTTPError(statusCode int, errorCode enums.ErrorCode, errorMessage enums.ErrorMessage) error {
	return &HTTPError{
		StatusCode:   statusCode,
		ErrorCode:    errorCode,
		ErrorMessage: errorMessage,
	}
}

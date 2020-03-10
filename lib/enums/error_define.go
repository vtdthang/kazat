package enums

// ErrorCode define all error codes
type ErrorCode int

//
const (
	UserNotFoundErrCode                  ErrorCode = 1000
	EmailAlreadyExistsErrCode            ErrorCode = 1001
	PasswordIsIncorrectErrCode           ErrorCode = 1002
	TokenIsExpiredErrCode                ErrorCode = 1003
	TokenIsMalformedOrInvalidErrCode     ErrorCode = 1004
	AuthorizationHeaderIsRequiredErrCode ErrorCode = 1005

	ServerErrCode ErrorCode = 9999
)

// ErrorMessage define all error messages
type ErrorMessage string

//
const (
	UserNotFoundErrMsg                   ErrorMessage = "User not found"
	EmailAlreadyExistsErrMsg             ErrorMessage = "Email already exists"
	PasswordIsIncorrectErrMsg            ErrorMessage = "Password is incorrect"
	TokenIsExpiredErrMsg                 ErrorMessage = "Token is expired"
	TokenIsMalformedOrInvalidErrMsg      ErrorMessage = "Token is malformed or invalid"
	AuthorizationHeaderIsRequiredErrCMsg ErrorMessage = "Authorization header is required"

	ServerErrMsg ErrorMessage = "Server is not responding..."
)

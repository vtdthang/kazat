package enums

// ErrorCode define all error codes
type ErrorCode int

//
const (
	UserNotFoundErrCode       ErrorCode = 1000
	EmailAlreadyExistsErrCode ErrorCode = 1001

	ServerErrCode ErrorCode = 9999
)

// ErrorMessage define all error messages
type ErrorMessage string

//
const (
	UserNotFoundErrMsg       ErrorMessage = "User not found"
	EmailAlreadyExistsErrMsg ErrorMessage = "Email already exists"
	ServerErrMsg             ErrorMessage = "Server is not responding..."
)

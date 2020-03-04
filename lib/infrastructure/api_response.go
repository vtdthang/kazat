package apiresponse

import (
	"encoding/json"
	"net/http"

	httpHeaderConstants "github.com/vtdthang/goapi/lib/constants/httpheader"
	mimeTypeConstants "github.com/vtdthang/goapi/lib/constants/mimetype"
)

// AsSuccessResponse return a success response
func AsSuccessResponse(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set(httpHeaderConstants.ContentType, mimeTypeConstants.ApplicationJSON)

	json.NewEncoder(w).Encode(body)
}

// AsErrorResponse return an error response
func AsErrorResponse(w http.ResponseWriter, statusCode int, errorCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	w.Header().Set(httpHeaderConstants.ContentType, mimeTypeConstants.ApplicationJSON)

	body := map[string]interface{}{
		"error_message": errorMessage,
		"error_code":    errorCode,
	}

	json.NewEncoder(w).Encode(body)
}

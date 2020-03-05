package apiresponse

import (
	"encoding/json"
	"net/http"

	httpHeaderConstants "github.com/vtdthang/goapi/lib/constants/httpheader"
	mimeTypeConstants "github.com/vtdthang/goapi/lib/constants/mimetype"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
)

// AsSuccessResponse return a success response
func AsSuccessResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set(httpHeaderConstants.ContentType, mimeTypeConstants.ApplicationJSON)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(body)
}

// AsErrorResponse return an error response
func AsErrorResponse(w http.ResponseWriter, err error) {
	switch errType := err.(type) {
	case *httperror.HTTPError:
		w.Header().Set(httpHeaderConstants.ContentType, mimeTypeConstants.ApplicationJSON)
		w.WriteHeader(errType.StatusCode)

		json.NewEncoder(w).Encode(err)
	default:
		w.Header().Set(httpHeaderConstants.ContentType, mimeTypeConstants.ApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)

		body := map[string]interface{}{
			"error_message": enums.ServerErrMsg,
			"error_code":    enums.ServerErrCode,
		}

		json.NewEncoder(w).Encode(body)
	}
}

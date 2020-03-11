package apiresponse

import (
	"encoding/json"
	"net/http"

	"github.com/vtdthang/goapi/lib/constants"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
)

// AsSuccessResponse return a success response
func AsSuccessResponse(w http.ResponseWriter, body interface{}) {
	w.Header().Set(constants.HTTPHeaderContentType, constants.MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(body)
}

// AsErrorResponse return an error response
func AsErrorResponse(w http.ResponseWriter, err error) {
	switch errType := err.(type) {
	case *httperror.HTTPError:
		w.Header().Set(constants.HTTPHeaderContentType, constants.MIMEApplicationJSON)
		w.WriteHeader(errType.StatusCode)

		json.NewEncoder(w).Encode(err)
	default:
		w.Header().Set(constants.HTTPHeaderContentType, constants.MIMEApplicationJSON)
		w.WriteHeader(http.StatusInternalServerError)

		body := map[string]interface{}{
			"error_message": enums.ServerErrMsg,
			"error_code":    enums.ServerErrCode,
		}

		json.NewEncoder(w).Encode(body)
	}
}

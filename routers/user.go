package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vtdthang/goapi/lib/constants"
)

// SetUsersRoutes is ...
func SetUsersRoutes(router *httprouter.Router) *httprouter.Router {
	router.GET("/users", getUsers)

	return router
}

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(constants.HTTPHeaderContentType, constants.MIMEApplicationJSON)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET Users"}`))
}

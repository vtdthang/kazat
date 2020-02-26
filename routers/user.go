package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SetUsersRoutes is ...
func SetUsersRoutes(router *httprouter.Router) *httprouter.Router {
	router.GET("/users", getUsers)

	return router
}

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET Users"}`))
}

package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// SetPostRoutes abc
func SetPostRoutes(router *httprouter.Router) *httprouter.Router {
	router.GET("/posts", getPosts)

	return router
}

func getPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET Posts"}`))
}

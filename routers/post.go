package routers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/vtdthang/goapi/lib/constants/httpheader"
	"github.com/vtdthang/goapi/lib/constants/mimetype"
)

// SetPostRoutes abc
func SetPostRoutes(router *httprouter.Router) *httprouter.Router {
	router.GET("/posts", getPosts)

	return router
}

func getPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(httpheader.ContentType, mimetype.ApplicationJSON)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "GET Posts"}`))
}

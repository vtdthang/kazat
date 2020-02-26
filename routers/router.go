package routers

import "github.com/julienschmidt/httprouter"

// InitRoutes is ...
func InitRoutes() *httprouter.Router {
	router := httprouter.New()
	router = SetPostRoutes(router)
	router = SetUsersRoutes(router)

	return router
}

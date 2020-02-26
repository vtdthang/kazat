package main

import (
	"log"
	"net/http"

	client "github.com/vtdthang/goapi/drivers/mongo"
	"github.com/vtdthang/goapi/routers"
)

func main() {

	client.ConnectSingleton()

	router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", router))
}

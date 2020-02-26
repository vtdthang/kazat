package main

import (
	"log"
	"net/http"

	"github.com/vtdthang/goapi/routers"
)

func main() {
	log.Println("OK")

	router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", router))
}

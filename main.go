package main

import (
	//"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	pgdatabase "github.com/vtdthang/goapi/drivers/pg"
	"github.com/vtdthang/goapi/routers"
)

func main() {
	connectToPostgres()

	router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", router))
}

func connectToPostgres() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	pgConnectionString := os.Getenv("PG_LOCAL_URL")

	pgdatabase.NewDB(pgConnectionString)

	fmt.Println("Successfully connected!")
}

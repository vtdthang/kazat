package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	pgdb "github.com/vtdthang/goapi/drivers/pg"
	"github.com/vtdthang/goapi/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	pgConnectionString := os.Getenv("PG_LOCAL_URL")
	db, err := pgdb.NewDB(pgConnectionString)
	if err != nil {
		fmt.Println("Cannot connect to Postgres!")
	}

	router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", router))
}

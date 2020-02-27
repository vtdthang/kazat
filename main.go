package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	psqlInfo := os.Getenv("PG_LOCAL_URL")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	pgdb "github.com/vtdthang/goapi/drivers/pg"
	userController "github.com/vtdthang/goapi/user/controller"
	userRepo "github.com/vtdthang/goapi/user/repository"
	userService "github.com/vtdthang/goapi/user/service"
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

	db.SetConnMaxLifetime(500)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(10)
	db.Stats()

	defer db.Close()

	router := httprouter.New()

	userRepo := userRepo.NewUserRepository(db)
	userService := userService.NewUserService(userRepo)
	userController := userController.NewUserController(userService)

	router.GET("/api/users/test", userController.FindByEmail)
	//router := routers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8081", router))
}

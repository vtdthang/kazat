package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	pgdb "github.com/vtdthang/goapi/drivers/pg"
	"github.com/vtdthang/goapi/lib/constants"
	"github.com/vtdthang/goapi/middlewares"
	userController "github.com/vtdthang/goapi/user/controller"
	userRepo "github.com/vtdthang/goapi/user/repository"
	userService "github.com/vtdthang/goapi/user/service"
)

func main() {
	// TEST HTML Template /////////////
	tmpl := template.Must(template.ParseFiles("layout.html"))
	////// END Test Html Template

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		panic("Error loading env file")
	}

	db, err := pgdb.NewDB(os.Getenv(constants.EnvPostgresURL))
	if err != nil {
		fmt.Println(err)
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

	router.GET("/", func(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		data := testWelcome{TextWelcome: "Welcome to Go REST API"}
		tmpl.Execute(w, data)
	})

	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/secured", middlewares.AuthorizeMiddleware(userController.Secured))
	//router := routers.InitRoutes()

	log.Fatal(http.ListenAndServe(":8081", router))
}

type testWelcome struct {
	TextWelcome string
}

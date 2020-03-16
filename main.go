package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

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

	// Serve static assets via the "static" directory
	router.ServeFiles("/static/*filepath", http.Dir("static"))

	// Method 2 ---> below
	dir, _ := os.Getwd()
	fmt.Println(dir)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", router)

	router.HandlerFunc("GET", "/", serveTemplate)

	router.POST("/api/users/register", userController.Register)
	router.POST("/api/users/login", userController.Login)
	router.GET("/api/users/secured", middlewares.AuthorizeMiddleware(userController.Secured))
	//router := routers.InitRoutes()

	log.Fatal(http.ListenAndServe(":8081", router))
}

// Render template
func serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}

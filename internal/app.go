package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	db     DataBaseService
}

func (a *App) Initialize(user string, password string, dbname string) {
	connectionString := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		user,
		password,
		dbname,
	)
	a.db = NewDB("postgres", connectionString)

	a.Router = mux.NewRouter()

	a.initializeRoutes()
}

func (a *App) Run(s string) {
	fmt.Println("Starting server on " + s)
	log.Fatal(http.ListenAndServe(s, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/courses", a.CreateCourse).Methods("POST")
}

func (a *App) CreateCourse(writer http.ResponseWriter, request *http.Request) {
}

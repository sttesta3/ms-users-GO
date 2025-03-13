package internal

import (
	"encoding/json"
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
	a.Router.Use(headers)
	a.Router.HandleFunc("/courses", a.CreateCourse).Methods("POST")
}

func headers(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}

func (a *App) insertCourse(c Course) error {
	err := a.db.InsertCourse(c)
	return err

}

func (a *App) CreateCourse(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var c Course
		err := json.NewDecoder(request.Body).Decode(&c)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = a.insertCourse(c)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(writer).Encode("Success")
	}
}

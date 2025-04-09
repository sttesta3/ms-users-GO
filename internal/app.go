package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Db     DataBaseService
}

func (a *App) Initialize(user string, password string, dbname string) {
	host := os.Getenv("DATABASE_HOST")
	if host != "" {
		connectionString := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s sslmode=disable",
			host,
			user,
			password,
			dbname,
		)
		a.Db = NewDB("postgres", connectionString)
	
		a.Router = mux.NewRouter()
	
		a.initializeRoutes()
	}
}

func (a *App) Run(s string) {
	fmt.Println("Starting server on " + s)
	log.Fatal(http.ListenAndServe(s, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.Use(headers)
	a.Router.HandleFunc("/courses", a.CreateCourse).Methods("POST")
	a.Router.HandleFunc("/courses", a.GetCourses).Methods("GET")
	a.Router.HandleFunc("/courses/{id}", a.GetCourse).Methods("GET")
	a.Router.HandleFunc("/courses/{id}", a.DeleteCourse).Methods("DELETE")
}

func (a *App) internalServerError(writer http.ResponseWriter) {
	writer.WriteHeader(500)
	errResponse := ErrorResponse{
		Status: 500,
		Title:  "Internal server error",
		Description:  "The server was unable to complete your request",
	}
	json.NewEncoder(writer).Encode(errResponse)
	return
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

func (a *App) insertCourse(c Course) (int, error) {
	return a.Db.InsertCourse(c)
}

func (a *App) getCourses() ([]Course, error) {
	return a.Db.GetCourses()
}

func (a *App) getCourse(id int) (Course, error) {
	return a.Db.GetCourse(id)
}

func (a *App) deleteCourse(id int) error {
	return a.Db.DeleteCourse(id)
}

func (a *App) CreateCourse(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var c Course
		err := json.NewDecoder(request.Body).Decode(&c)
		if c.Title == "" || c.Description == "" {

			writer.WriteHeader(400)
			errResponse := ErrorResponse{
				Status:      400,
				Title:       "No se indicó la información adecuada.",
				Description: "No se especificó un título o descripción para el curso.",
			}
			json.NewEncoder(writer).Encode(errResponse)
			return

		}

		if err != nil {
			a.internalServerError(writer)
			fmt.Println(err.Error())
			return
		}

		id, err := a.insertCourse(c)
		if err != nil {
			a.internalServerError(writer)
			fmt.Println(err.Error())
			return
		}
		c.Id = &id
		writer.WriteHeader(201)
		json.NewEncoder(writer).Encode(map[string]Course{
			"data": c,
		})
	}
}

func (a *App) GetCourses(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		courses, err := a.getCourses()
		if err != nil {
			a.internalServerError(writer)
			fmt.Println(err.Error())
			return
		}
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(map[string][]Course{
			"data": courses,
		})
	}
}

func (a *App) GetCourse(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		courseIdStr := mux.Vars(request)["id"]
		courseId, err := strconv.Atoi(courseIdStr)
		if err != nil {
			http.Error(writer, err.Error(), 400)
			return
		}
		course, err := a.getCourse(courseId)
		if err != nil {
			errResponse := ErrorResponse{
				Status:      404,
				Title:       "Course not found",
				Description: "No se encontró el curso especificado",
			}
			writer.WriteHeader(404)
			json.NewEncoder(writer).Encode(errResponse)
			return
		}
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(map[string]Course{
			"data": course,
		})
	}
}

func (a *App) DeleteCourse(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		courseIdStr := mux.Vars(request)["id"]
		courseId, err := strconv.Atoi(courseIdStr)
		if err != nil {
			http.Error(writer, err.Error(), 400)
			return
		}
		err = a.deleteCourse(courseId)
		if err != nil {
			errResponse := ErrorResponse{
				Status:      404,
				Title:       "Course not found",
				Description: "No se encontró el curso especificado",
			}
			writer.WriteHeader(404)
			json.NewEncoder(writer).Encode(errResponse)
			return
		}
		writer.WriteHeader(204)
		json.NewEncoder(writer).Encode("Course deleted succesfully")
	}
}

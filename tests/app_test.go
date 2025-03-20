package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"ing2-tp1/internal"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
)

type MockDB struct{}

func (mockDb *MockDB) InsertCourse(c internal.Course) (int, error) {
	return 1, nil
}

func (mockDb *MockDB) GetCourses() ([]internal.Course, error) {
	return make([]internal.Course, 0), nil
}

func (mockDb *MockDB) GetCourse(num int) (internal.Course, error) {
	if num == 100 {
		return internal.Course{}, fmt.Errorf("error")
	}
	return internal.Course{}, nil
}

func (mockDb *MockDB) DeleteCourse(num int) error {
	if num == 1 {
		return nil
	}
	return fmt.Errorf("error")
}

func TestPostCoursesOK(t *testing.T) {
	course := internal.Course{
		Title:       "a",
		Description: "asdasd",
	}
	jsonCourse, _ := json.Marshal(course)
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonCourse))
	w := httptest.NewRecorder()

	app.CreateCourse(w, r)
	if w.Code != 201 {
		t.Errorf("Result was incorrect")
	}
}

func TestPostCoursesSinInformacion(t *testing.T) {
	course := internal.Course{}
	jsonCourse, _ := json.Marshal(course)
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonCourse))
	w := httptest.NewRecorder()

	app.CreateCourse(w, r)
	if w.Code != 400 {
		t.Errorf("Result was incorrect")
	}
}

func TestGetCourses(t *testing.T) {
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("GET", "/courses", nil)
	w := httptest.NewRecorder()

	app.GetCourses(w, r)
	if w.Code != 200 {
		t.Errorf("Result was incorrect")
	}
}

func TestGetCourse(t *testing.T) {
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("GET", "/courses/1", nil)
	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	app.GetCourse(w, r)
	if w.Code != 200 {
		t.Errorf("Result was incorrect")
	}
}

func TestGetCourseNotFound(t *testing.T) {
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("GET", "/courses/100", nil)
	vars := map[string]string{
		"id": "100",
	}

	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	app.GetCourse(w, r)
	if w.Code != 404 {
		t.Errorf("Result was incorrect")
	}
}

func TestDeleteCourse(t *testing.T) {
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("DELETE", "/courses/1", nil)
	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	app.DeleteCourse(w, r)
	if w.Code != 204 {
		t.Errorf("Result was incorrect")
	}
}

func TestDeleteCourseMissingCourse(t *testing.T) {
	app := internal.App{}
	app.Initialize(
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PWD"),
		os.Getenv("POSTGRES_DB"),
	)

	app.Db = &MockDB{}

	r, _ := http.NewRequest("DELETE", "/courses/10", nil)
	vars := map[string]string{
		"id": "10",
	}

	r = mux.SetURLVars(r, vars)

	w := httptest.NewRecorder()

	app.DeleteCourse(w, r)
	if w.Code != 404 {
		t.Errorf("Result was incorrect")
	}
}

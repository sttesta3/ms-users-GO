package tests

import (
	"bytes"
	"encoding/json"
	"ing2-tp1/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDB struct{}

func (mockDb *MockDB) InsertCourse(c internal.Course) (int, error) {
	return 1, nil
}

func TestPostCoursesOK(t *testing.T) {
	course := internal.Course{
		Title:       "a",
		Description: "asdasd",
	}
	jsonCourse, _ := json.Marshal(course)
	app := internal.App{}
	app.Initialize(
		"postgres",
		"1234",
		"ingsoft2")

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
		"postgres",
		"1234",
		"ingsoft2")

	app.Db = &MockDB{}

	r, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(jsonCourse))
	w := httptest.NewRecorder()

	app.CreateCourse(w, r)
	if w.Code != 400 {
		t.Errorf("Result was incorrect")
	}
}

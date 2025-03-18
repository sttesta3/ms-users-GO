package internal

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type DataBaseService interface {
	InsertCourse(c Course) (int, error)
	GetCourses() ([]Course, error)
	GetCourse(id int) (Course, error)
}

type PostgresService struct {
	*sql.DB
}

func (self *PostgresService) CreateCourse() int {
	return 0
}

func NewDB(dbDriver string, dbSource string) DataBaseService {
	newDb, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		panic(err)
	}

	if dbDriver == "postgres" {
		return &PostgresService{DB: newDb}
	}
	return nil
}

func (self *PostgresService) InsertCourse(c Course) (int, error) {
	query := "INSERT INTO courses (title, description) VALUES($1, $2);"
	querySelect := "SELECT id FROM courses WHERE title = $1 and description = $2"
	_, err := self.Exec(query, c.Title, c.Description)
	row := self.QueryRow(querySelect, c.Title, c.Description)
	var id int
	row.Scan(&id)
	return id, err
}

func (self *PostgresService) GetCourses() ([]Course, error) {
	query := "SELECT * FROM courses"
	rows, err := self.Query(query)
	if err != nil {
		return nil, err
	}
	courses := make([]Course, 0)
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.Id, &course.Title, &course.Description); err != nil {
			log.Fatal(err)
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (self *PostgresService) GetCourse(id int) (Course, error) {
	query := "SELECT * FROM courses where id = $1"
	row := self.QueryRow(query, id)
	var course Course
	var idSql sql.NullInt64
	err := row.Scan(&idSql, &course.Title, &course.Description)
	if idSql.Valid {
		intId := int(idSql.Int64)
		course.Id = &intId
	}
	return course, err
}

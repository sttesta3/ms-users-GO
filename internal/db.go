package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBaseService interface {
	InsertCourse(c Course) (int, error)
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

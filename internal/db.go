package internal

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DataBaseService interface {
	CreateCourse() int
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

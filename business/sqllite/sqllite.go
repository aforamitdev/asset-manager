package sqllite

import (
	"changeme/foundation/validate"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	ErrorDBNotFound = "Database not found"
)

func NewDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		return nil, validate.NewAppError(err, "failed to ")
	}

	return db, nil

}

package mocks

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/prest/adapters/postgres"
)

func MockDB(driver string) (mock sqlmock.Sqlmock, err error) {
	var db *sql.DB
	db, mock, err = sqlmock.New()
	if err != nil {
		return
	}
	postgres.AddDatabaseToPool("knab", sqlx.NewDb(db, driver))
	postgres.SetDatabase("knab")
	return
}

package db

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb/v2"
)

type DBProvider struct {
	db *sql.DB
}

func NewDBProvider() *DBProvider {
	return &DBProvider{}
}

func (p *DBProvider) DBConnect(dbname string) (*sql.DB, func() error) {
	db, err := sql.Open("duckdb", dbname)
	if err != nil {
		panic(err)
	}

	p.db = db

	return db, db.Close
}

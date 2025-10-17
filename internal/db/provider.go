package db

import (
	"database/sql"

	_ "github.com/marcboeker/go-duckdb/v2"
)

type DBProvider struct {
	conn *sql.DB
}

func NewDBProvider() *DBProvider {
	return &DBProvider{}
}

func (p *DBProvider) DBConnect() (*sql.DB, func() error, error) {
	conn, err := sql.Open("duckdb", ":memory:")
	if err != nil {
		return nil, conn.Close, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, conn.Close, err
	}
	p.conn = conn

	return conn, conn.Close, nil
}

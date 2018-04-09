package main

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq"
	// _ "github.com/go-sql-driver/mysql"
)

// NewDBConn creates a new database client and returns the database version
func NewDBConn(dbURL string, dbType string) (*sql.DB, string, error) {
	var err error
	var db *sql.DB

	switch dbType {
	case "postgres":
		db, err = sql.Open("postgres", dbURL)
	case "mysql":
		db, err = sql.Open("mysql", dbURL)
	default:
		return nil, "", errors.New("Unknown Database type")
	}
	if err != nil {
		return nil, "", err
	}

	// Verify db connection
	err = db.Ping()
	if err != nil {
		return nil, "", err
	}
	var version string
	err = db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return nil, "", err
	}

	return db, version, nil
}

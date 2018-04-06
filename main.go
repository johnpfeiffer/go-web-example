package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	// _ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := NewDBConn("dbname=mydb host=127.0.0.1 user=myuser password=mypassword port=5432 sslmode=disable", "postgres")
	exitIfError(err)
	notes, err := getNotes(db)
	fmt.Println(notes)
	exitIfError(err)
	fmt.Println("done")
}

func exitIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// NewDBConn creates a new database client
func NewDBConn(dbURL string, dbType string) (*sql.DB, error) {
	var err error
	var db *sql.DB

	switch dbType {
	case "postgres":
		db, err = sql.Open("postgres", dbURL)
	case "mysql":
		db, err = sql.Open("mysql", dbURL)
	default:
		return nil, errors.New("Unknown Database type")
	}
	if err != nil {
		return nil, err
	}

	// Verify db connection
	err = db.Ping()
	if err != nil {
		return db, err
	}
	var DBversion string
	err = db.QueryRow("SELECT version()").Scan(&DBversion)
	if err != nil {
		return db, err
	}

	return db, nil
}

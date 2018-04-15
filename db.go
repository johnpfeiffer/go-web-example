package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	// _ "github.com/go-sql-driver/mysql"
)

func getDBCredentials() string {
	dbhost := getEnvOrDefault("DB_HOST", "127.0.0.1")
	dbport := getEnvOrDefault("DB_PORT", "5432")
	dbssl := getEnvOrDefault("DB_SSL", "disable")
	dbuser := getEnvOrDefault("DB_USERNAME", "myuser")
	dbpassword := getEnvOrDefault("DB_PASSWORD", "mypassword")
	dbname := getEnvOrDefault("DB_NAME", "mydb")
	dbConnString := fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s",
		dbname, dbhost, dbuser, dbpassword, dbport, dbssl)
	return dbConnString
}

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

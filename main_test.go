package main

// integration tests that modify a database

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	// TODO: remove the flag since the environment variable is possibly more portable/universal
	databaseTest = flag.Bool("databaseTest", false, "run database integration tests")
	testDB       *sql.DB
)

// TestMain allows for custom setup and running instead of main
// https://golang.org/pkg/testing/#hdr-Main , http://cs-guy.com/blog/2015/01/test-main/ , https://www.philosophicalhacker.com/post/integration-tests-in-go/
func TestMain(m *testing.M) {
	flag.Parse()
	// if *databaseTest {
	if isIntegrationTest() {
		setupTestDB()
	}
	exitCode := m.Run()
	os.Exit(exitCode)
}

func isIntegrationTest() bool {
	val, ok := os.LookupEnv("TEST_INTEGRATION")
	if ok && val == "true" {
		return true
	}
	return false
}

func setupTestDB() {
	dbhost := getEnvOrDefault("TEST_DB_HOST", "127.0.0.1")
	dbport := getEnvOrDefault("TEST_DB_PORT", "5432")
	dbssl := getEnvOrDefault("TEST_DB_SSL", "disable")
	dbuser := getEnvOrDefault("TEST_DB_USERNAME", "myuser")
	dbpassword := getEnvOrDefault("TEST_DB_PASSWORD", "mypassword")
	dbname := getEnvOrDefault("TEST_DB_NAME", "mydb")
	dbConnString := fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s",
		dbname, dbhost, dbuser, dbpassword, dbport, dbssl)
	db, dbVersion, err := NewDBConn(dbConnString, "postgres")
	exitIfError(err)
	fmt.Println("test database is setup:", dbVersion)
	testDB = db
	if testDB == nil {
		log.Fatal(err)
	}
	if dbVersion == "" {
		log.Fatal("ERROR: unable to get the test database version")
	}
	err = testDB.Ping()
	if err != nil {
		log.Fatal("ERROR: unable to ping the test database")
	}
}

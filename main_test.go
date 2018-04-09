package main

// integration tests that modify a database

import (
	"fmt"
	"log"
	"os"
	"testing"
)

// TestMain allows for custom setup and running instead of main  https://golang.org/pkg/testing/#hdr-Main
func TestMain(m *testing.M) {

	// call flag.Parse() here if TestMain uses flags
	dbhost := getEnvOrDefault("TEST_DB_HOST", "127.0.0.1")
	dbport := getEnvOrDefault("TEST_DB_PORT", "5432")
	dbssl := getEnvOrDefault("TEST_DB_SSL", "disable")
	dbuser := getEnvOrDefault("TEST_DB_USERNAME", "myuser")
	dbpassword := getEnvOrDefault("TEST_DB_PASSWORD", "mypassword")
	dbname := getEnvOrDefault("TEST_DB_NAME", "mydb")
	dbConnString := fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s",
		dbname, dbhost, dbuser, dbpassword, dbport, dbssl)
	db, dbVersion, err := NewDBConn(dbConnString, "postgres")

	if err != nil {

	}
	if db == nil {

	}
	if dbVersion == "" {

	}

	testText := "a test note"
	testNote := Note{Text: testText}
	preNotes, err := getNotes(db)
	if err != nil {
		log.Fatal(err)
	}
	if len(preNotes) != 0 {
		log.Fatal("No notes should exist yet in the test")
	}
	noteID, err := createNote(db, testNote)
	if err != nil {
		log.Fatal(err)
	}
	if noteID != 1 {
		log.Fatal("A single test note should exist with id 1, instead received id", noteID)
	}

	postNotes, err := getNotes(db)
	if len(postNotes) != 1 {
		log.Fatal("A single test note should exist")
	}
	if postNotes[0].Text != testText {
		log.Fatal("A test note text should be", testText, "but instead received", postNotes[0].Text)
	}

	exitCode := m.Run()

	//clearDB()
	os.Exit(exitCode)
}

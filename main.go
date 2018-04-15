package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// TODO: logging

	dbConnString := getDBCredentials()
	db, dbVersion, err := NewDBConn(dbConnString, "postgres")
	exitIfError(err)
	fmt.Printf("Database Version: %s \n", dbVersion)

	notes, err := getNotes(db)
	fmt.Println(notes)
	exitIfError(err)

	r := router()
	http.Handle("/", r)
	// TODO: env variable to set the port
	// TODO: more custom control over handling timeouts
	http.ListenAndServe(":8080", r)
}

func exitIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	result := defaultValue
	val, ok := os.LookupEnv(key)
	if ok {
		result = val
	}
	return result
}

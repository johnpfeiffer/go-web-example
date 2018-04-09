package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	dbhost := getEnvOrDefault("DB_HOST", "127.0.0.1")
	dbport := getEnvOrDefault("DB_PORT", "5432")
	dbssl := getEnvOrDefault("DB_SSL", "disable")
	dbuser := getEnvOrDefault("DB_USERNAME", "myuser")
	dbpassword := getEnvOrDefault("DB_PASSWORD", "mypassword")
	dbname := getEnvOrDefault("DB_NAME", "mydb")
	dbConnString := fmt.Sprintf("dbname=%s host=%s user=%s password=%s port=%s sslmode=%s",
		dbname, dbhost, dbuser, dbpassword, dbport, dbssl)
	db, dbVersion, err := NewDBConn(dbConnString, "postgres")
	exitIfError(err)
	fmt.Printf("Database Version: %s \n", dbVersion)

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

func getEnvOrDefault(key, defaultValue string) string {
	result := defaultValue
	val, ok := os.LookupEnv(key)
	if ok {
		result = val
	}
	return result
}

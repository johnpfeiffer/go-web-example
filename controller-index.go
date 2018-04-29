package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var defaultIndexResponse = `<html><body>index web page</body></html>`

// IndexHandler sends the Index page HTTP response
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, defaultIndexResponse)
}

// NoteHandlerGET returns the current Notes
func NoteHandlerGET(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: logging wrapper
	log.Println("received GET")
	notes, err := getNotes(db)
	// TODO: better error handling
	exitIfError(err)
	s := fmt.Sprintf("%v", notes)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, s)
}

// NoteHandlerPOST inserts a note
func NoteHandlerPOST(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("received POST")
	defer r.Body.Close()
	responseData, err := ioutil.ReadAll(r.Body)
	exitIfError(err)

	n := Note{Text: string(responseData)}
	noteID, err := createNote(db, n)
	log.Printf("inserted note with id: %d", noteID)
}

// TODO NotesHandler , /notes/{date}
// 	vars := mux.Vars(r)
// fmt.Fprintf(w, "you have requested notes from date: %v\n", vars["date"])
// TODO: totally unsafe for XSS etc.

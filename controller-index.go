package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// IndexHandler sends the Index page HTTP response
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body>index web page</body></html>")
}

// NoteHandlerGET returns the current Notes
func NoteHandlerGET(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// TODO: logging wrapper
	log.Println("received GET")
	notes, err := getNotes(db)
	// TODO: better error handling
	exitIfError(err)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if len(notes) < 1 {
		fmt.Fprintf(w, "{}")
		return
	}
	theJSON, err := json.MarshalIndent(notes, "", "  ")
	exitIfError(err)
	fmt.Fprintf(w, string(theJSON))
}

// NoteHandlerPOST inserts a note
func NoteHandlerPOST(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	log.Println("received POST")
	defer r.Body.Close()
	responseData, err := ioutil.ReadAll(r.Body)
	// TODO: better error handling
	exitIfError(err)

	var n Note
	err = json.Unmarshal(responseData, &n)
	exitIfError(err)

	noteID, err := createNote(db, n)
	log.Printf("inserted note with id: %d", noteID)
}

// TODO NotesHandler , /notes/{date}
// 	vars := mux.Vars(r)
// fmt.Fprintf(w, "you have requested notes from date: %v\n", vars["date"])
// TODO: totally unsafe for XSS etc.

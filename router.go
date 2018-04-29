package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

func router(db *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	r.HandleFunc("/note",
		func(w http.ResponseWriter, r *http.Request) {
			NoteHandlerGET(w, r, db)
		}).Methods("GET")
	r.HandleFunc("/note",
		func(w http.ResponseWriter, r *http.Request) {
			NoteHandlerPOST(w, r, db)
		}).Methods("POST")

	return r
}

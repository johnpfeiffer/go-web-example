package main

import (
	"fmt"
	"net/http"
)

// IndexHandler sends the Index page HTTP response
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<html><body>index web page</body></html>")
}

// TODO NotesHandler , /notes/{date}
// 	vars := mux.Vars(r)
// fmt.Fprintf(w, "you have requested notes from date: %v\n", vars["date"])
// TODO: totally unsafe for XSS etc.

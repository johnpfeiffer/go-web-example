package main

import (
	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")
	return r
}

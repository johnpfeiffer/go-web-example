Title: Leveraging Gorilla Mux for a Webserver with Postgres
Date: 2018-02-26 19:19
Tags: go, golang, web, http, gorilla mux, postgres, pg, psql

## Prerequisites

    # install docker in order to trivially run postgres locally
    go get github.com/lib/pq github.com/gorilla/mux
    sudo ./start-db.sh
    sudo ./create_tables.sh
    sudo ./psql.sh
    \l
    \dt+
    \d+
    \d+ TABLENAME 


## Basic Code

The application has different parts in order to separate concerns:

*start-db.sh and create-tables.sh are helpers to setup a local dev environment using docker*

1. main.go starts the application
2. notes.go represents an example "model" (like MVC or <https://en.wikipedia.org/wiki/Data_access_object>)

## Basic Execution

`go run main.go note.go`

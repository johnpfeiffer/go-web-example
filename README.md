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
2. db.go abstracts the persistence layer connection
3. notes.go represents an example "model" (like MVC or <https://en.wikipedia.org/wiki/Data_access_object>)

## Basic Execution

`go run main.go db.go note.go`

## Basic Testing

Unit testing should not require external dependencies, the -short command can still skip "long" unit tests

`go test` or `go test -v`

### Integration

Since the integration tests expect to actively use a real database there is an environment variable that tells the system how to initialize

> Otherwise the integration tests are skipped, no database needed!

`TEST_INTEGRATION=true go test -v` (helpfully wrapped in integration_tests.sh)

Environment variable and the default:

    "TEST_DB_HOST", "127.0.0.1"
    "TEST_DB_PORT", "5432"
    "TEST_DB_SSL", "disable"
    "TEST_DB_USERNAME", "myuser"
    "TEST_DB_PASSWORD", "mypassword"
    "TEST_DB_NAME", "mydb"




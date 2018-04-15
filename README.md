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


## Running it

`./run.sh` (*which just executes something like go run main.go db.go note.go*)


### Manually checking the web server
`curl localhost:8080`
> <html><body>index web page</body></html>

`curl -I -X GET localhost:8080`

    HTTP/1.1 200 OK
    Date: Sun, 15 Apr 2018 17:48:47 GMT
    Content-Length: 40
    Content-Type: text/html; charset=utf-8


### Manually checking the database
`sudo ./psql.sh`


## Basic Testing

Unit testing should not require external dependencies, *the -short command can still skip "long" unit tests*

`go test` or `go test -v`

### Integration Testing

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



## Code Structure

The application has different parts in order to separate concerns:

*start-db.sh and create-tables.sh are helpers to setup a local dev environment using docker*

0. tables.sql is the database schema
1. main.go initializes and starts the application
2. db.go abstracts the persistence layer connection
3. notes.go represents an example "model" (like MVC or <https://en.wikipedia.org/wiki/Data_access_object>)
4. router.go manage the routes to the web application
5. controller-.go are the definitions of handlers for each route




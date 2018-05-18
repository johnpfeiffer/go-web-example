#!/bin/bash

sudo docker run -it --rm --link some-postgres:postgres  -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --dbname mydb

# IF no default database was created...
# sudo docker run -it --rm --link some-postgres:postgres  -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser


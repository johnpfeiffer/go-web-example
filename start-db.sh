#!/bin/bash
# https://hub.docker.com/_/postgres/
set -x
sudo docker run --rm -it -p 0.0.0.0:5432:5432 --name some-postgres -e POSTGRES_PASSWORD=mypassword -e POSTGRES_USER=myuser -e POSTGRES_DB=mydb  postgres:alpine

# alternative without watching it run and a more restrictive access to the Postgres service
# sudo docker run --rm --name some-postgres --publish 127.0.0.1:5432:5432  -e POSTGRES_PASSWORD=mypassword -e POSTGRES_USER=myuser -e POSTGRES_DB=mydb postgres:alpine

# ALTERNATIVELY create a default database of myuser (shrug)
# sudo docker run --rm -it -p 0.0.0.0:5432:5432 --name some-postgres -e POSTGRES_PASSWORD=mypassword -e POSTGRES_USER=myuser  postgres:alpine 


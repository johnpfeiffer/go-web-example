#!/bin/bash
sudo docker run --rm --link some-postgres:postgres -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --dbname mydb --command '\l+'

# IF the mydb database was not created already
# sudo docker run --rm --link some-postgres:postgres  -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --command 'create table mydb;'

# sudo docker run --rm --link some-postgres:postgres -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --dbname mydb --command 'CREATE TABLE IF NOT EXISTS example( id SERIAL PRIMARY KEY );'

# for some reason interactive mode is required ;)
sudo docker run -i --rm --link some-postgres:postgres -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --dbname mydb < tables.sql

sudo docker run --rm --link some-postgres:postgres  -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --dbname mydb --command '\dt+'
# sudo docker run --rm --link some-postgres:postgres  -e PGPASSWORD=mypassword postgres:alpine psql --host postgres --username myuser --dbname mydb --command '\d+ notes'

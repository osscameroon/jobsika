#!/bin/bash

#This should not be run as a standalone script but only from the docker-compose

set -u

ENV_FILE=/api/.docker-env

#install goose
if [[ -z "$(which goose)" ]]; then
	go install github.com/pressly/goose/v3/cmd/goose@v3.15.1
fi

#install psql
if [[ -z "$(which psql)" ]]; then
	apt-get update
	apt-get -y install postgresql-client
fi

cd /api/scripts
./goose-up.sh $ENV_FILE

export $(grep -v '#.*' $ENV_FILE | xargs)
PG_STRING="host=$POSTGRES_HOST dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

cd /api/scripts/setup-postgres
while read -r file; do
	psql "$PG_STRING" -f "$file"
done < ./sql_file_ordered.txt

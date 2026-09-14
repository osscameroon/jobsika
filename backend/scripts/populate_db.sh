#!/bin/bash

ENV_FILE=../.postgres-env
if [ $# -ge 1 ]; then
	ENV_FILE=$1
fi

#install goose
if [[ -z "$(which goose)" ]]; then
	go install github.com/pressly/goose/v3/cmd/goose@v3.15.1
fi

#install psql
if [[ -z "$(which psql)" ]]; then
	apt-get update
	apt-get -y install postgresql-client
fi

./goose-up.sh $ENV_FILE

export $(grep -v '#.*' $ENV_FILE | xargs)
PG_STRING="host=$POSTGRES_HOST dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

cd ./setup-postgres
while read -r file; do
	psql "$PG_STRING" -f "$file"
done < ./sql_file_ordered.txt

#!/bin/bash

ENV_FILE=../.postgres-env
if [ $# -ge 1 ]; then
	ENV_FILE=$1
fi

export $(grep -v '#.*' $ENV_FILE | xargs)
PG_STRING="host=$POSTGRES_HOST dbname=$POSTGRES_DB user=$POSTGRES_USER password=$POSTGRES_PASSWORD sslmode=disable"

cd ../migrations/
goose postgres "$PG_STRING" reset

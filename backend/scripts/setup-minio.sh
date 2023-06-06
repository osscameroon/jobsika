#!/bin/bash

#This should not be run as a standalone script but only from the docker-compose

set -u

ENV_FILE=/api/.docker-env

#install minio client
if [[ -z "$(which mc)" ]]; then
	go install github.com/minio/mc@latest
fi

export $(grep -v '#.*' $ENV_FILE | xargs)

mc alias set local http://minio:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc admin user svcacct add --access-key "$MINIO_ACCESS_KEY" --secret-key "$MINIO_SECRET_KEY" local $MINIO_ROOT_USER

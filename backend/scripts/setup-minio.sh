#!/bin/bash

#This should not be run as a standalone script but only from the docker-compose

set -u

ENV_FILE=/api/.docker-env

#install minio client
if [[ -z "$(which mc)" ]]; then
	wget -O mc https://dl.min.io/client/mc/release/linux-amd64/mc
	chmod +x mc
	mv mc /usr/local/bin/mc
fi

export $(grep -v '#.*' $ENV_FILE | xargs)

mc alias set local http://minio:9000 $MINIO_ROOT_USER $MINIO_ROOT_PASSWORD
mc mb local/$S3_BUCKET
mc admin user svcacct add --access-key "$S3_ACCESS_KEY" --secret-key "$S3_SECRET_KEY" local $MINIO_ROOT_USER

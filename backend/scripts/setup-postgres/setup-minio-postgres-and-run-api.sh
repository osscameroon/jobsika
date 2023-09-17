#!/bin/bash

#file used only with docker
/api/scripts/wait-for-it/wait-for-it.sh minio:9000 -t 120 -- /api/scripts/setup-minio.sh
/api/scripts/setup-postgres/setup-postgres.sh
make -C /api/ run

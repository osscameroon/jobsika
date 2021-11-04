#!/bin/bash

#file used only with docker
/api/scripts/setup-postgres/setup-postgres.sh
make -C /api/ e2etest

# image used by docker-compose to apply migrations and load fixtures
# it ships psql, bash and goose so setup-postgres.sh does not install anything at runtime
FROM golang:1.24-alpine AS goose

RUN go install github.com/pressly/goose/v3/cmd/goose@v3.26.0

FROM postgres:14.1-alpine

RUN apk add --no-cache bash

COPY --from=goose /go/bin/goose /usr/local/bin/goose

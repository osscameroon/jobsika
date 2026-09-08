FROM golang:alpine AS base

# update the ca-certificates
# more on https://github.com/google/go-github/issues/1049
RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

#copies go.sum and go.mod if exists
COPY go.* ./
RUN go mod download

COPY . .

RUN go build -o api .

WORKDIR /dist

RUN cp /build/api .

# build image with the binary
FROM scratch

# copy the ca-certificate.crt from the build stage
# more on https://github.com/google/go-github/issues/1049#issuecomment-1023836581
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=base /dist/api /

# expose api port
EXPOSE 7001

ENTRYPOINT ["/api"]

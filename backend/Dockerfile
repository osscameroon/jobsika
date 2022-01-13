FROM golang:alpine AS base

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

COPY --from=base /dist/api /

# expose api port
EXPOSE 7000

ENTRYPOINT ["/api"]

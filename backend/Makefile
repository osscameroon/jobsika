APPNAME=jobsika
E2ETEST_DEPS=./e2etests/node_modules

.DEFAULT_GOAL := help

define install_goose
	go install github.com/pressly/goose/v3/cmd/goose@latest
endef


## test: run tests on cmd and pkg files.
.PHONY: test
test: vet fmt
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

## build: build application binary.
.PHONY: build
build:
	go build -o $(APPNAME)

## run: run the api
.PHONY: run
run:
	ENVIRONMENT=development go run .

## run4test: run the api in test environment (fixes timeouts)
.PHONY: run4test
run4test:
	ENVIRONMENT=development TEST=true go run .

## e2etest: run end to end tests against local api
## : run the api using `ENVIRONMENT=development TEST=true make run` or `make run4test` to avoid timeouts
.PHONY: e2etest
e2etest: $(E2ETEST_DEPS)
	-cd ./e2etests/ && API_HOST="http://localhost:7000/" npm test
	@make reset-postgres > /dev/null

## e2etest-compose: run end to end tests in the docker-compose.test.yaml
.PHONY: e2etest-compose
e2etest-compose: $(E2ETEST_DEPS)
	cd ./e2etests/ && npm test

$(E2ETEST_DEPS):
	cd ./e2etests && npm install

## docker-e2etest: run e2etests in a docker compose
docker-e2etest: $(E2ETEST_DEPS)
	-rm -rf postgres-data
	docker-compose -f docker-compose.test.yml up --abort-on-container-exit --exit-code-from e2etests

## docker-build: build the api docker image
.PHONY: docker-build
docker-build:
	docker build -t jobsika .

## docker-run: run the api docker container
.PHONY: docker-run
docker-run:
	docker run -p 7000:7000 --env-file .docker-env jobsika

.PHONY: install_goose
install_goose:
	$(call install_goose)

## goose_up: run apply database schema migration
.PHONY: install_goose
goose_up:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./goose-up.sh

## goose_down: undo a migration
goose_down:
.PHONY: goose_down
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./goose-down.sh

## goose_reset: reset migrations
.PHONY: goose_reset
goose_reset:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./goose-reset.sh

## start-api: starts the api and its dependencies in a docker container
start-api:
	docker-compose up

## stop-api: stops the api and its dependencies in a docker container
stop-api:
	docker-compose down && \
	rm -rf postgres-data

## start-postgres: starts postgres sql, setup and populate yotas database
start-postgres:
	docker-compose up start-postgres
	docker-compose logs -f

## connect-postgres: connect to postgres sql
connect-postgres:
	cd ./scripts/ && ./connect_to_postgres.sh

## clean-postgres: delete every tables stored in the yotas database
clean-postgres:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./goose-reset.sh

## populate-postgres: populate the database with new data
populate-postgres:
ifeq (, $(shell which goose))
	$(call install_goose)
endif
	cd ./scripts/ && ./populate_db.sh

## reset-postgres: clean and populate the data base
reset-postgres:
	make clean-postgres && make populate-postgres

## stop-postgres: remove local postgres-data and teardown postgres
stop-postgres:
	docker-compose down && \
	rm -rf postgres-data

.PHONY:check-swagger
check-swagger:
	which swagger || (go install github.com/go-swagger/go-swagger/cmd/swagger@latest)

swagger.yaml: check-swagger
	swagger generate spec -o ./swagger.yaml

## serve-swagger: open the api documentation
.PHONY:serve-swagger
serve-swagger: swagger.yaml check-swagger
	swagger serve -F=swagger swagger.yaml

## install-deps: install main dependencies
install-deps:
	@echo "installing golangci-lint..."
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.8
	@echo "golangci installed!"
	@echo "installing errcheck..."
	go install github.com/kisielk/errcheck@latest
	@echo "errcheck installed!"
	@echo "installing goose..."
	go install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "goose installed!"
	@echo "installing swagger..."
	go install github.com/go-swagger/go-swagger/cmd/swagger@latest
	@echo "swagger installed!"


check-lint:
ifeq (, $(shell which golangci-lint))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.23.8
endif
ifeq (, $(shell which errcheck))
	go install github.com/kisielk/errcheck@latest
endif

## lint: run linters over the entire code base
.PHONY: lint
lint: check-lint
	golangci-lint run ./... --timeout 15m0s
	errcheck -exclude ./.golangci-errcheck-exclude.txt ./...

## install-hooks: install hooks
.PHONY: install-hooks
install-hooks:
	ln -s $(PWD)/githooks/pre-push .git/hooks/pre-push

## clean: remove releases
.PHONY: clean
clean:
	rm -rf $(APPNAME)

all: help
.PHONY: help
help: Makefile
	@echo " Choose a command..."
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

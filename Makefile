.DEFAULT_GOAL := help

# use the docker compose v2 plugin, fallback to the legacy docker-compose binary
DOCKER_COMPOSE ?= $(shell docker compose version > /dev/null 2>&1 && echo 'docker compose' || echo docker-compose)
# same go version as the CI, the pinned golangci-lint does not support newer ones
GO_LINT_IMAGE=golang:1.16
# frontend commands run in the node container of the frontend-dev service
FRONTEND_RUN=$(DOCKER_COMPOSE) run --rm --no-deps frontend-dev sh -c
FRONTEND_DEPS=yarn install --frozen-lockfile --silent

## build: build every docker images (db setup, api and frontend)
.PHONY: build
build:
	$(DOCKER_COMPOSE) build

## build-backend: build the api binary (requires go)
.PHONY: build-backend
build-backend:
	make -C backend build

## build-frontend: generate the static frontend in ./frontend/dist (in docker)
.PHONY: build-frontend
build-frontend:
	$(FRONTEND_RUN) "$(FRONTEND_DEPS) && BASE_URL=http://localhost:7000 yarn generate"

## lint: lint the backend and the frontend
.PHONY: lint
lint: lint-backend lint-frontend

## lint-backend: lint the backend (in docker)
.PHONY: lint-backend
lint-backend:
	docker run --rm -v $(PWD)/backend:/app -v jobsika-go-lint-cache:/go -w /app $(GO_LINT_IMAGE) make lint

## lint-frontend: lint the frontend (in docker)
.PHONY: lint-frontend
lint-frontend:
	$(FRONTEND_RUN) "$(FRONTEND_DEPS) && yarn lint"

## lintfix-frontend: fix the frontend lint errors (in docker)
.PHONY: lintfix-frontend
lintfix-frontend:
	$(FRONTEND_RUN) "$(FRONTEND_DEPS) && yarn lintfix"

## start: run the whole stack (api on port 7000, frontend on port 3000)
.PHONY: start
start:
	$(DOCKER_COMPOSE) up --build

## dev: run the api and the frontend dev server with hot reload (port 3000)
.PHONY: dev
dev:
	$(DOCKER_COMPOSE) --profile dev up --build frontend-dev

## stop: stop every containers
.PHONY: stop
stop:
	$(DOCKER_COMPOSE) --profile dev down

## clean: stop every containers and delete the database and the volumes
.PHONY: clean
clean:
	$(DOCKER_COMPOSE) --profile dev down -v

## logs: follow the logs of every containers
.PHONY: logs
logs:
	$(DOCKER_COMPOSE) --profile dev logs -f

## test: run the backend unit tests (requires go)
.PHONY: test
test:
	make -C backend test

## e2etest: run the backend end to end tests (in docker)
.PHONY: e2etest
e2etest:
	make -C backend docker-e2etest

## install-hooks: install hooks
.PHONY: install-hooks
install-hooks:
	ln -s $(PWD)/githooks/pre-push .git/hooks/pre-push

all: help
.PHONY: help
help: Makefile
	@echo " Choose a command..."
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

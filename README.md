# Jobsika
JobSika is a website that allows people to anonymously rate their companies and share their salaries, in order to create a more level playing field between employers and employees. People can contribute their salary information to the website and help to change the current narrative around salary information.

### HOW TO CONTRIBUTE TO JobSika

First, fork the project from the GitHub repository. To do this, click the "Fork" button in the top-right corner of the repository page and follow the instructions.

#### Quick start with Docker

The only requirements are GNU make and [docker](https://docs.docker.com/get-docker/) with the compose plugin (v2.24+).
Everything is driven by the root `Makefile`, run `make help` to list every command.

- Run the whole stack (database, migrations + fixtures, api and frontend): `make start`
  - api: http://localhost:7000/
  - frontend: http://localhost:3000/
- Run the api and the frontend dev server (hot reload) instead of the built frontend: `make dev`
- Follow the logs: `make logs`
- Stop everything: `make stop`
- Stop everything and delete the database: `make clean`

Payments are disabled unless you copy `backend/.opencollective-env-example` to `backend/.opencollective-env` and fill it.

#### Build, lint and test

| Command | Description |
|---|---|
| `make build` | build every docker images (db setup, api and frontend) |
| `make build-backend` | build the api binary in `./backend/jobsika` (requires go) |
| `make build-frontend` | generate the static frontend in `./frontend/dist` (in docker) |
| `make lint` | lint the backend and the frontend (in docker) |
| `make lint-backend` / `make lint-frontend` | lint only one side (in docker) |
| `make lintfix-frontend` | fix the frontend lint errors (in docker) |
| `make test` | run the backend unit tests (requires go) |
| `make e2etest` | run the backend end to end tests against a dockerized api (requires node to install the tests dependencies) |
| `make install-hooks` | install the git pre-push hook |

The GitHub workflows run these same commands, so if they pass locally they should pass on the CI.

#### Run without Docker

You will need:

- Golang version 1.16 and above
- GNU make
- Node.js v16 (see `frontend/.nvmrc`, e.g. `nvm use`), newer versions are not supported by the Nuxt 2 dependencies
- docker, to run the database
- [swagger](https://goswagger.io/install.html), to serve the api documentation

Backend, from the `./backend` folder (`make help` lists every command):

- Run the database with migrations and fixtures: `make start-postgres`
- Run the api: `make run`
- Build the api: `make build && ./jobsika`
- Serve the api documentation: `make serve-swagger`
- Stop the database and delete its data: `make stop-postgres`

Frontend, from the `./frontend` folder:

- Install the dependencies: `yarn install`
- Run the app: `yarn dev`, you can then access it through http://localhost:3000/
- The api it talks to is set with the `BASE_URL` environment variable (see `frontend/README.md`)

#### Others docs to check

**How to create a new table:** [how-to-add-a-new-table.md](./docs/how-to-add-a-new-table.md)

Aaaand that's it!
Thank you for taking the time to read this document. Following the steps outlined in this document will help you set up and run the application locally. We hope this document has been helpful and that it makes it easier for you to contribute and help make jobsika better.

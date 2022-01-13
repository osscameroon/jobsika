# camerdevs
Camerdevs, is dev cameroonian developer platform.

### Features
The platform includes the functionnalities bellow:
- A Catalogue cameroonian developers
- Their availabilies
- The developers salaries broken down by companies
- And later some job offers

### Development

#### Dependencies
To run locally, this api require few dependencies

- Golang version 1.16 or later
- GNU make
- docker with docker-compose

#### How to run
Once you have installed the required depencies, you can now run the application
following the patterns bellow

#### Backend
To run the backend you should jump to the `./backend` folder
`cd ./backend`

##### Run

First run the database:
`make start-postgres`

Then run the api:
`make run`

##### Build and run
`make run && ./camerdevs`

#### Build and run with docker
`make docker-build && make docker-run`

#### Server api documentation
`make serve-swagger`

#### How to run the api with the DB
`make start-api`

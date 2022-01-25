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

- Golang version 1.16
- GNU make
- docker with docker-compose
- [swagger](https://goswagger.io/install.html)

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

#### Frontend

To run the frontend you should jump to the ./frontend folder using the command cd ./frontend

First install the frontend's dependencies:

`npm install` or
`yarn install`

Then build the frontend app:

`npm run build` or
`yarn run build`

And then run the app:

`npm run start` or
`yarn run start`

Now that the application is running, you can access it through http://localhost:3000/
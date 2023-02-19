# Jobsika
JobSika is a website that allows people to anonymously rate their companies and share their salaries, in order to create a more level playing field between employers and employees. People can contribute their salary information to the website and help to change the current narrative around salary information.

### Development

#### Dependencies
To run locally, this api require few dependencies

- Golang version 1.16
- GNU make
- docker with docker-compose
- [swagger](https://goswagger.io/install.html)

#### How to run
Once you have installed the required dependencies, you can now run the application
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
`make run && ./jobsika`

#### Build and run with docker
`make docker-build && make docker-run`

#### Server api documentation
`make serve-swagger`

#### How to run the api with the DB
`make start-api`

#### Others docs to check

**How to create a new table:** [how-to-add-a-new-table.md](./docs/how-to-add-a-new-table.md)

#### Frontend

To run the frontend you should jump to the ./frontend folder using the command cd ./frontend and run the following commands:

`npm install` or
`yarn install` to install the dependencies.

`npm run dev` or
`yarn dev` to run the app.

Now that the application is running, you can access it through http://localhost:3000/

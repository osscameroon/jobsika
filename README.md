# Jobsika
JobSika is a website that allows people to anonymously rate their companies and share their salaries, in order to create a more level playing field between employers and employees. People can contribute their salary information to the website and help to change the current narrative around salary information.

### Development

#### Dependencies
To run the application locally, you should:
1. Fork the project from the GitHub repository. To do this, click the "Fork" button in the top-right corner of the repository page and follow the instructions.
2. Install the required dependencies:

- Golang version 1.16
- GNU make
- docker with docker-compose
- [swagger](https://goswagger.io/install.html)

#### Run the application following the steps below:
- Jump to the ./backend folder: `cd ./backend`
- Run the database: `make start-postgres`
- Run the api: `make run`
- Build and run: `make run && ./jobsika`
- Build and run with docker: `make docker-build && make docker-run`
- Serve the api documentation: `make serve-swagger`
- Run the api with the DB: `make start-api`
#### Others docs to check

**How to create a new table:** [how-to-add-a-new-table.md](./docs/how-to-add-a-new-table.md)

#### Jump to the ./frontend folder: cd ./frontend
- Install the dependencies: `npm install` or `yarn install`
- Run the app: `npm run dev` or `yarn dev`
- Now that the application is running, you can access it through http://localhost:3000/

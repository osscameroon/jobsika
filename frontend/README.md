## Frontend

### Dependencies

Make sure you are using Node.js v16 (see `.nvmrc`). Newer versions are not supported by the Nuxt 2 dependencies.
You can also run the dev server in docker from the root of the repository: `make dev`. You can use [nvm](https://github.com/nvm-sh/nvm) to manage multiple Node.js versions.

### Run the frontend

Before running the frontend, you may want to create a `.env` file,
copy the content of `.env.example` and paste it in the newly created `.env`.
You can now edit the `BASE_URL` which indicates where your frontend will be pointing to.

```bash
BASE_URL=https://api.jobsika.stage.osscameroon.com
```

Note: it is possible that the backend does not yet contain the data you want to display.
In this case, you can run the frontend against the local backend.

To run the backend, run `make start` from the root of the repository. You can read the `../README.md` for more information.

Once the backend is running, edit your `.env` file with `BASE_URL=http://localhost:7000`.

#### Build machinery

```bash
# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev

# build for production and launch server
$ yarn build

# generate static project
$ yarn generate

# launch the server's production build
$ yarn start
```

For detailed explanation on how things work, check out the [documentation](https://nuxtjs.org).

If you find any issues while running the application please open an [issue](https://github.com/osscameroon/jobsika/issues/new) or ask a question on our [telegram group](https://t.me/+UpKZh_KXTaTx7JD7).

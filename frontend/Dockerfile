FROM node:12.22.9-buster-slim
WORKDIR /usr/src/app
COPY . /usr/src/app
RUN npm install && npm run generate

FROM flashspys/nginx-static
RUN apk update && apk upgrade
COPY --from=0 /usr/src/app/dist /static

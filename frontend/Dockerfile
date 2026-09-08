FROM node:12.22.9-buster-slim

ARG baseURL="http://localhost:7001"
ENV BASE_URL=$baseURL

WORKDIR /usr/src/app
COPY . .
RUN yarn install
RUN yarn generate

FROM flashspys/nginx-static
RUN apk update && apk upgrade
COPY --from=0 /usr/src/app/dist /static

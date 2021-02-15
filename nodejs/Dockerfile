FROM node:15.8.0-alpine3.13

WORKDIR /usr/src/

RUN apk add --no-cache bash

COPY ./react .

WORKDIR /usr/src/app

RUN npm install

CMD [ "yarn", "start" ]
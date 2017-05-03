FROM node:alpine

COPY . /usr/src/app

WORKDIR /usr/src/app

RUN yarn install && yarn build

RUN mv build/* /go/src/github.com/gotoolkit/goreact/web

VOLUME ["/go/src/github.com/gotoolkit/goreact/web"]
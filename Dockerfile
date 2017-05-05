FROM golang:alpine

RUN apk add --no-cache git curl && curl https://glide.sh/get | sh

COPY . /go/src/github.com/gotoolkit/goreact

RUN mv web/build/* web/ && rm -rf web/build/

WORKDIR /go/src/github.com/gotoolkit/goreact

RUN glide install && go install

ENTRYPOINT ["goreact"]


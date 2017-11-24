FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go install app

ENTRYPOINT /go/bin/app

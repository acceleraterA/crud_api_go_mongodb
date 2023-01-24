FROM golang:1.16-alpine as builder

WORKDIR /
COPY ./src .

RUN go build -o /crud-api

ENTRYPOINT ["/crud-api"]
FROM golang:1.19-alpine AS build
LABEL MAINTAINER github.com/arizon-dread
WORKDIR /usr/src/go-desperado
EXPOSE 8080
ENV GIN_MODE=release
COPY go.mod go.sum ./
RUN apk update && apk add --no-cache git 
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/go-desperado ./...
CMD [ "go-desperado" ]
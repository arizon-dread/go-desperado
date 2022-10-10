FROM golang:1.19-alpine AS build
LABEL MAINTAINER github.com/arizon-dread
ENV GIN_MODE=release
ENV PORT=8080

COPY businesslayer /usr/local/go/src/github.com/arizon-dread/go-desperado/businesslayer
COPY models /usr/local/go/src/github.com/arizon-dread/go-desperado/models
COPY main.go go.mod go.sum /usr/local/go/src/github.com/arizon-dread/go-desperado/

RUN apk update && apk add --no-cache git
WORKDIR /usr/local/go/src/github.com/arizon-dread/go-desperado
RUN go build ./...


FROM alpine:3.16 AS final
WORKDIR /go/
COPY --from=build /usr/local/go/src/github.com/arizon-dread/go-desperado/go-desperado /go/bin/
EXPOSE ${PORT}

ENTRYPOINT [ "/go/bin/go-desperado" ]
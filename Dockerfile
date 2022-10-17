FROM golang:1.19-alpine AS build
LABEL MAINTAINER github.com/arizon-dread

#COPY businesslayer /usr/local/go/src/github.com/arizon-dread/go-desperado/businesslayer
#COPY models /usr/local/go/src/github.com/arizon-dread/go-desperado/models
#COPY main.go go.mod go.sum /usr/local/go/src/github.com/arizon-dread/go-desperado/
WORKDIR /usr/local/go/src/github.com/arizon-dread/go-desperado
COPY . . 
RUN apk update && apk add --no-cache git 
RUN go build -v -o /usr/local/bin/go-desperado ./...


FROM golang:1.19-alpine AS final
WORKDIR /go/bin
ENV GIN_MODE=release
RUN apk add --no-cache libc6-compat musl-dev
COPY --from=build /usr/local/bin/go-desperado /go/bin/go-desperado
EXPOSE 8080

ENTRYPOINT [ "./go-desperado" ]

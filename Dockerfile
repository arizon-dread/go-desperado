FROM golang:1.19-alpine AS build
LABEL MAINTAINER github.com/arizon-dread

#COPY businesslayer /usr/local/go/src/github.com/arizon-dread/go-desperado/businesslayer
#COPY models /usr/local/go/src/github.com/arizon-dread/go-desperado/models
#COPY main.go go.mod go.sum /usr/local/go/src/github.com/arizon-dread/go-desperado/
WORKDIR /usr/local/go/src/github.com/arizon-dread/go-desperado
COPY . . 
RUN apk update && apk add --no-cache git gcc musl-dev util-linux-dev
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags "-linkmode external -extldflags '-static' -s -w" ./...


FROM golang:1.19-alpine AS final
WORKDIR /go/bin
ENV GIN_MODE=release
ENV PORT=8080
RUN apk add --no-cache libc6-compat musl-dev
COPY --from=build /usr/local/go/src/github.com/arizon-dread/go-desperado/go-desperado /go/bin/go-desperado
EXPOSE ${PORT}

ENTRYPOINT [ "./go-desperado" ]

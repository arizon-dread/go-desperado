FROM golang:alpine AS build
LABEL MAINTAINER github.com/arizon-dread
ENV GIN_MODE=release
ENV PORT=8080

COPY businesslayer/ models/ main.go go.mod go.sum /go/src/

RUN apk update && apk add --no-cache git
WORKDIR /go/src
RUN go get ./...

RUN go build .

FROM golang:alpine AS final

WORKDIR /go/

COPY --from=build /go/src/go-desperado /go/

EXPOSE ${PORT}

ENTRYPOINT [ "./go-desperado" ]
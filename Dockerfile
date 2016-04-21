FROM golang:1.6
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/meshblu-connector-dependency-manager
COPY . /go/src/github.com/octoblu/meshblu-connector-dependency-manager

RUN env CGO_ENABLED=0 go build -o meshblu-connector-dependency-manager -a -ldflags '-s' .

CMD ["./meshblu-connector-dependency-manager"]

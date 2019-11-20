FROM golang:1.13

RUN apt-get update

ADD . /go/src/userland
WORKDIR /go/src/userland

CMD go get ./... && go run main.go
EXPOSE 8080

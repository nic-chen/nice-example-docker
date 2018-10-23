FROM golang:latest

RUN go get github.com/nic-chen/nice

COPY src/ $GOPATH/src/nice-example/

WORKDIR $GOPATH/src/nice-example

RUN go build

ENTRYPOINT ["./nice-example"]
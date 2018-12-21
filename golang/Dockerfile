FROM golang:latest

RUN go get github.com/nic-chen/nice
#RUN go get github.com/golang/protobuf
RUN go get github.com/grpc-ecosystem/go-grpc-middleware
RUN go get github.com/grpc-ecosystem/go-grpc-prometheus
RUN go get github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc
RUN go get github.com/opentracing/opentracing-go
RUN go get github.com/uber/jaeger-client-go/config
RUN go get go.etcd.io/etcd/clientv3
RUN go get go.uber.org/zap

COPY src/ $GOPATH/src/nice-example/

WORKDIR $GOPATH/src/nice-example

RUN make build

ENTRYPOINT ["cmd/nice-test", "all"]
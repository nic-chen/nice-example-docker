package controller

import (
	"time"
	"context"
	
	"../../config"

	"github.com/nic-chen/nice"
	"github.com/nic-chen/nice/micro/tracing"

	"google.golang.org/grpc"
	"github.com/nic-chen/nice/micro/registry/etcdv3"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
)

type Controller struct {
	Name string
}

type JsonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func New(name string) *Controller {
	ctl := &Controller{
		Name: name,
	}

	return ctl;
}

func RenderJson(c *nice.Context, code int, message string, data interface{}) {
	ret := new(JsonResp)
	ret.Code = code
	ret.Message = message
	if data != nil {
		ret.Data = data
	}
	c.JSON(200, ret)
}

func newSrvDialer(serviceName string) *grpc.ClientConn {
	r := etcdv3.NewResolver(serviceName)
	b := grpc.RoundRobin(r)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	tracer, err := tracing.Init(config.CliName, config.JaegerAddr)
	conn, err := grpc.DialContext(ctx, config.NamingAddr, grpc.WithInsecure(), grpc.WithBalancer(b), grpc.WithBlock(), grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer)))
	cancel()
	if err != nil {
		panic(err)
	}

	return conn
}


package framework

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"google.golang.org/grpc"
)

func NewGrpcServer() *grpc.Server {
	srv := grpc.NewServer()
	return srv
}

func NewHttpServer() *ghttp.Server {
	srv := ghttp.GetServer()
	return srv
}

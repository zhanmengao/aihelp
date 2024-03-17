package book

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"google.golang.org/grpc"
)

func Init(gSrv *grpc.Server, hSrv *ghttp.Server) {
	srv := &bookService{}
	pb.RegisterAisrvBookServer(gSrv, srv)
}

type bookService struct {
}

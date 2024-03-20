package cos

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"google.golang.org/grpc"
)

func Init(gSrv *grpc.Server, hSrv *ghttp.Server) {
	srv := &cosService{}
	pb.RegisterAisrvCosServer(gSrv, srv)
}

type cosService struct {
}

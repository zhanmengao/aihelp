package writer

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"google.golang.org/grpc"
)

func Init(gSrv *grpc.Server, hSrv *ghttp.Server) {
	srv := &writeService{}
	pb.RegisterApsrvWriterServer(gSrv, srv)
}

type writeService struct {
}

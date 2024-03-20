package book

import (
	"context"
	"fmt"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"github.com/zhanmengao/aihelp/module/framework"
	"github.com/zhanmengao/aihelp/module/gpt"
)

// ReadBook 一条新会话，将想要读的书发过去
func (b bookService) ReadBook(ctx context.Context, req *pb.ReadBookReq) (rsp *pb.ReadBookRsp, err error) {
	rsp = &pb.ReadBookRsp{}
	//创建一个SessionID
	msg := fmt.Sprintf("《%s》这本书讲述了什么？写下你读这本书的感受和书评", req.BookName)
	retMsg, err := gpt.NewGpt(ctx, framework.GetUID(ctx), framework.NewID()).SendMessage(ctx, "", msg)
	if err != nil {
		return
	}
	rsp.Message = retMsg

	return
}

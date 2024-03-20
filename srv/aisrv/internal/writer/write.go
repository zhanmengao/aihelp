package writer

import (
	"context"
	"fmt"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"github.com/zhanmengao/aihelp/module/framework"
	"github.com/zhanmengao/aihelp/module/gpt"
)

func (w writeService) Write(ctx context.Context, req *pb.WriteReq) (rsp *pb.WriteRsp, err error) {
	rsp = &pb.WriteRsp{}
	topic := ""
	switch req.Write {
	case pb.WriteWorkSummary:
		topic = fmt.Sprintf("帮我写一篇工作总结，我的工作如下：%s", req.Desc)
	case pb.WriteTopic:
		topic = fmt.Sprintf("围绕 [%s] 帮我写一篇文章", req.Desc)
	}
	style := ""
	switch req.Style {
	case pb.WriteStyleSolemnly:
		style = "使用严肃认真的口吻"
	case pb.WriteStyleHumorous:
		style = "使用幽默的口吻"
	}
	retMsg, err := gpt.NewGpt(ctx, framework.GetUID(ctx), framework.NewID()).SendMessage(ctx, "", fmt.Sprintf("%s，%s", topic, style))
	if err != nil {
		return
	}
	rsp.Message = retMsg
	return
}

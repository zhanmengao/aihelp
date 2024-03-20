package cos

import (
	"context"
	"fmt"
	"github.com/zhanmengao/aihelp/manifest/protobuf/pb"
	"github.com/zhanmengao/aihelp/module/framework"
	"github.com/zhanmengao/aihelp/module/gpt"
)

func (c cosService) CosPlay(ctx context.Context, req *pb.CosPlayReq) (rsp *pb.CosPlayRsp, err error) {
	rsp = &pb.CosPlayRsp{}
	//创建一个SessionID
	role := "朋友"
	switch req.RoleType {
	case pb.CosPlayBoyFriend:
		role = "男朋友"
	case pb.CosPlayCat:
		role = "猫"
	case pb.CosPlayGirlFriend:
		role = "女朋友"
	}
	msg := fmt.Sprintf("需要你扮演我的%s和我对话。你是一个这样的人：%s", role, req.RoleDesc)
	retMsg, err := gpt.NewGpt(ctx, framework.GetUID(ctx), framework.NewID()).SendMessage(ctx, "", msg)
	if err != nil {
		return
	}
	rsp.Message = retMsg
	return
}

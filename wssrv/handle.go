package wssrv

import (
	"context"
	"github.com/zhanmengao/gf/proto/go/pb"
)

type IHandler interface {
	//HandleInit
	//  @Description: 传递sessionID，业务调用FLS接口
	//  @param cli 客户端类封装
	//  @param sessionID SessionID
	//  @return uid 返回UID
	//  @return err 返回错误，出错时连接被关闭
	HandleInit(ctx context.Context, cli *Client, sessionKey string, rcv []byte) (uid, aesKey string, err error)

	//HandlePacket
	//  @Description: pkt过滤器
	//  @param cli 客户端类封装
	//  @param pkt 消息包
	//  @return ok 是否继续转发
	//
	HandlePacket(ctx context.Context, cli *Client, pkt *pb.Packet) (err error)

	//HandleClose
	//  @Description: 关闭前的回调，可以发送遗言
	//  @param ctx：上下文
	//  @param cli：客户端类的封装
	//
	HandleClose(ctx context.Context, cli *Client, err error)
}

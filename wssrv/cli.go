package wssrv

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/zhanmengao/gf/proto/go/pb"
)

type Client struct {
	request *ghttp.Request
	conn    *ghttp.WebSocket
	srv     *WsSrv
	xRealIP string
}

func NewClient(ws *WsSrv, request *ghttp.Request, conn *ghttp.WebSocket) *Client {
	c := &Client{
		request: request,
		conn:    conn,
		srv:     ws,
		xRealIP: "",
	}
	ws.clientMap.Store(c.GetKey(), c)
	return c
}

func (c *Client) GetKey() string {
	return c.xRealIP
}

func (c *Client) Run() {
	go func() {
		defer func() {
			c.srv.clientMap.Delete(c.GetKey())
			c.conn.Close()
		}()
		for {
			ctx := gctx.New()
			_, body, err := c.conn.ReadMessage()
			if err != nil {
				glog.Warningf(ctx, "%s connect close [%s]", c.GetKey(), err)
				return
			}
			//body 解析后回调上层
			pkt := &pb.Packet{}
			if err = pkt.Unmarshal(body); err != nil {
				glog.Warningf(ctx, "%s connect close [%s]", c.GetKey(), err)
				return
			}
			if err = c.srv.handle.HandlePacket(ctx, c, pkt); err != nil {
				glog.Warningf(ctx, "%s connect close [%s]", c.GetKey(), err)
				return
			}
		}
	}()
}

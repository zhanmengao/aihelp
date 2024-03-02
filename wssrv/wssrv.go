package wssrv

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"sync"
)

// WsSrv 微信小程序入口
type WsSrv struct {
	clientMap sync.Map
	srv       *ghttp.Server
	handle 	  IHandler
}

func NewServer() *WsSrv {
	s := g.Server()
	s.BindHandler("/wss", func(r *ghttp.Request) {
		ws, err := r.WebSocket()
		if err != nil {
			panic(err)
		}
		for {

		}
	})
}

func

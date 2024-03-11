package v35

import (
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/sashabaranov/go-openai"
	"github.com/zhanmengao/aihelp/config"
	"github.com/zhanmengao/aihelp/global"
	"github.com/zhanmengao/aihelp/manifest/protobuf/gendb"
	"github.com/zhanmengao/aihelp/module/gpt/client"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	global.PikaDB = gendb.NewPikaDBCache()
	config.GptConfig.V3ApiKey = "sk-WYd9zVPgx0ywV7CW634906Dc3b204aD3Bf6070E3Ef3426Cc"
	config.GptConfig.V3Host = "https://one.aiskt.com"
	m.Run()
}

func TestGpt3(t *testing.T) {
	gpt := TGpt35{
		UID:        "999",
		SessionKey: "999",
	}
	g.Client().Use(func(c *gclient.Client, r *http.Request) (rsp *gclient.Response, err error) {
		glog.Infof(gctx.New(), "midware")
		return nil, gerror.New("111")
	})
	err := gpt.SendMessage(gctx.New(), "我是一个程序员", "Steam小游戏创新方向")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGpt(t *testing.T) {
	cli := client.NewV3Client()

	ctx := gctx.New()
	req := openai.CompletionRequest{
		Model:       openai.GPT3TextDavinci003, // 选择的模型
		MaxTokens:   2048,
		N:           1,
		Stop:        nil,
		Temperature: 0.5,
		Prompt:      "写一篇《时间管理》", //要问的问题

	}

	resp, err := cli.CreateCompletion(ctx, req) // 发起接口调用
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(resp.Choices[0].Text) // 读取返回结果
}

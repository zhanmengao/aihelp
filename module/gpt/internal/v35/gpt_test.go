package v35

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/zhanmengao/aihelp/config"
	"github.com/zhanmengao/aihelp/global"
	"github.com/zhanmengao/aihelp/manifest/protobuf/gendb"
	"testing"
)

func TestMain(m *testing.M) {
	global.PikaDB = gendb.NewPikaDBCache()
	config.GptConfig.V3ApiKey = "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"
	config.GptConfig.V3Host = "https://one.aiskt.com"
	m.Run()
}

func TestGpt3(t *testing.T) {
	gpt := TGpt35{
		UID:        "999",
		SessionKey: "999",
	}
	err := gpt.SendMessage(gctx.New(), "我是一个程序员", "Steam小游戏创新方向")
	if err != nil {
		t.Fatal(err)
	}
}

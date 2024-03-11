package client

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sashabaranov/go-openai"
	"github.com/zhanmengao/aihelp/config"
	_ "github.com/zhanmengao/aihelp/global"
	"net/http"
	"reflect"
	"unsafe"
)

func NewV3Client() *openai.Client {
	cfg := openai.DefaultConfig(config.GptConfig.V3ApiKey)
	cfg.BaseURL = config.GptConfig.V3Host
	hcli := (*http.Client)(unsafe.Pointer(reflect.ValueOf(g.Client()).Pointer()))
	cfg.HTTPClient = hcli
	cli := openai.NewClientWithConfig(cfg)
	return cli
}

func NewV4Client() *openai.Client {
	cfg := openai.DefaultConfig(config.GptConfig.V4ApiKey)
	cfg.BaseURL = config.GptConfig.V3Host
	cli := openai.NewClientWithConfig(cfg)
	return cli
}

package client

import (
	"github.com/sashabaranov/go-openai"
	"github.com/zhanmengao/aihelp/config"
	_ "github.com/zhanmengao/aihelp/global"
)

func NewV3Client() *openai.Client {
	cfg := openai.DefaultConfig(config.GptConfig.V3ApiKey)
	cfg.BaseURL = config.GptConfig.V3Host
	cli := openai.NewClientWithConfig(cfg)
	return cli
}

func NewV4Client() *openai.Client {
	cfg := openai.DefaultConfig(config.GptConfig.V4ApiKey)
	cfg.BaseURL = config.GptConfig.V3Host
	cli := openai.NewClientWithConfig(cfg)
	return cli
}

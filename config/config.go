package config

import (
	"context"
	"github.com/zhanmengao/gf/config/icfg"
	"github.com/zhanmengao/gf/config/icfg/cfgtp"
)

type TGptConfig struct {
	ApiKey    string `yaml:"api_key"`
	HttpProxy string `yaml:"http_proxy"`
	AllProxy  string `yaml:"all_proxy"`
}

var (
	GptConfig TGptConfig
)

func Init(configFile string) (err error) {
	yCfg, err := icfg.NewConfig(context.Background(), configFile, cfgtp.YAML)
	if err != nil {
		return
	}
	err = yCfg.Get("gpt", &GptConfig)
	if err != nil {
		return
	}
	return
}

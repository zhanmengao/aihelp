package config

import (
	"encoding/json"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

type TGptConfig struct {
	ApiKey    string `yaml:"api_key" json:"api_key"`
	HttpProxy string `yaml:"http_proxy" json:"http_proxy"`
	AllProxy  string `yaml:"all_proxy" json:"all_proxy"`
}

var (
	GptConfig TGptConfig
)

func Init(configFile string) (err error) {
	yCfg := gcfg.Instance()
	yCfg.GetAdapter().(*gcfg.AdapterFile).SetFileName(configFile)

	cfg, err := yCfg.Get(gctx.New(), "gpt")
	if err != nil {
		return
	}
	if err = json.Unmarshal(cfg.Bytes(), &GptConfig); err != nil {
		return
	}
	return
}

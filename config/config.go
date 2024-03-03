package config

import (
	"encoding/json"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfsnotify"
	"github.com/gogf/gf/v2/os/glog"
)

type TGptConfig struct {
	ApiKey    string `yaml:"api_key" json:"api_key"`
	HttpProxy string `yaml:"http_proxy" json:"http_proxy"`
	AllProxy  string `yaml:"all_proxy" json:"all_proxy"`
}

var (
	GptConfig TGptConfig
	yCfg      *gcfg.Config
)

func Init(configFile string) (err error) {
	yCfg = gcfg.Instance()
	yCfg.GetAdapter().(*gcfg.AdapterFile).SetFileName(configFile)
	if err = reloadConfig(); err != nil {
		return
	}
	_, err = gfsnotify.Add(configFile, func(s *gfsnotify.Event) {
		//执行reload
		err = reloadConfig()
		if err != nil {
			glog.Fatalf(gctx.New(), "reload config error %s", err)
		}
	})
	if err != nil {
		return
	}
	return
}

func reloadConfig() (err error) {
	cfg, err := yCfg.Get(gctx.New(), "gpt")
	if err != nil {
		return
	}
	if err = json.Unmarshal(cfg.Bytes(), &GptConfig); err != nil {
		return
	}
	return
}

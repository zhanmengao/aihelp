package global

import "flag"

var (
	ConfigFile string
)

func Flags() {
	flag.StringVar(&ConfigFile, "c", "", "配置文件")
	flag.Parse()
}

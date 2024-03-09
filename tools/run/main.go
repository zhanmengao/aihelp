package main

import (
	"github.com/gogf/gf/v2/os/glog"
	"github.com/spf13/cobra"
	"github.com/zhanmengao/aihelp/tools/dbtool"
)

var rootCmd = &cobra.Command{
	Use:     "ai",
	Short:   "ai tool",
	Long:    `ai tool`,
	Version: "0.0.0",
}

func init() {
	rootCmd.AddCommand(dbtool.DBTool)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Error(rootCmd.Context(), err.Error())
	}
}

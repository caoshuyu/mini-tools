package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "mini-tools [OPTIONS] [COMMANDS]",
	Short: "mini tools 小工具合集",
	Long:  "mini tools 小工具合集。\n",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("数据参数 -h 查看相关帮助文档")
	},
}

// Execute 执行命令行信息
func Execute() {
	if err := useCmd.Execute(); err != nil {
		panic(err)
	}
}
